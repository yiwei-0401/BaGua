package controllers

import (
	"BaGua/common"
	"database/sql"
	gua64 "BaGua/models"
	yao368 "BaGua/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type Res struct {
	Code      int  `json:"code"`
	Status 	  string  `json:"status"`
	Data 	  interface{} `json:"data"`
}

type QiGuaController struct {
	beego.Controller
}

func (c *QiGuaController) QiGua(){
	var dayan = 49
	var yao string
	var bianyao string
	ben_gua := gua64.Ben_gua{}
	bian_gua := gua64.Bian_gua{}
	for i := 1; i < 7; i++ {
		//经过三变得到代表爻的数字
		yaoNum := common.Change(dayan, 1)
		yaoNum = yaoNum / 4
		//0表示阴爻，1表示阳爻
		switch yaoNum {
		case 6: //老阴
			yao = "0"
			bianyao = "1"
			ben_gua.BianYaos = append(ben_gua.BianYaos, i) //老阴为变卦
		case 7: //少阳
			bianyao = "1"
			yao = "1"
		case 8: //少阴
			bianyao = "0"
			yao = "0"
		case 9: //老阳
			yao = "1"
			bianyao = "0"
			ben_gua.BianYaos = append(ben_gua.BianYaos, i) //老阳变阴
		}
		//用6位二进制表示64卦
		ben_gua.BenguaCode.WriteString(yao)
		bian_gua.BianguaCode.WriteString(bianyao)
	}

	//连接数据库，从数据库中取出字符串代表的卦的行数据
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/zhouyi?charset=utf8");
	if err != nil {
		fmt.Printf("connect mysql fail ! [%s]", err)
	}
	benguaRows, err := db.Query("select id, name,alias,location,form,code,guaci,guaci_fanyi, yao_1,yao_2,yao_3,yao_4,yao_5,yao_6 from gua64 where code = " + ben_gua.BenguaCode.String());
	if err != nil {
		fmt.Printf("select fail [%s]", err)
	}
	//本卦集合
	ben_map := make(map[string]string)
	for benguaRows.Next() {
		err = benguaRows.Scan(&ben_gua.Id, &ben_gua.Name, &ben_gua.Alias, &ben_gua.Location, &ben_gua.Form, &ben_gua.Code, &ben_gua.Guaci, &ben_gua.GuaciFanyi, &ben_gua.Yao_1, &ben_gua.Yao_2, &ben_gua.Yao_3, &ben_gua.Yao_4, &ben_gua.Yao_5, &ben_gua.Yao_6)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		ben_map["name"] = ben_gua.Name
		ben_map["code"] = ben_gua.Code
		ben_map["alias"] = ben_gua.Alias
		ben_map["location"] = ben_gua.Location
		ben_map["form"] = ben_gua.Form
		ben_map["1"] = ben_gua.Yao_1
		ben_map["2"] = ben_gua.Yao_2
		ben_map["3"] = ben_gua.Yao_3
		ben_map["4"] = ben_gua.Yao_4
		ben_map["5"] = ben_gua.Yao_5
		ben_map["6"] = ben_gua.Yao_6
	}

	//变卦集合
	bian_map := make(map[string]string)
	bianguaRows, err := db.Query("select id, name,alias,location,form,code,guaci,guaci_fanyi, yao_1,yao_2,yao_3,yao_4,yao_5,yao_6 from gua64 where code = " + bian_gua.BianguaCode.String());

	if err != nil {
		fmt.Printf("select fail [%s]", err)
	}
	for bianguaRows.Next() {
		err = bianguaRows.Scan(&bian_gua.Id, &bian_gua.Name, &bian_gua.Alias, &bian_gua.Location, &bian_gua.Form, &bian_gua.Code, &bian_gua.Guaci, &bian_gua.GuaciFanyi, &bian_gua.Yao_1, &bian_gua.Yao_2, &bian_gua.Yao_3, &bian_gua.Yao_4, &bian_gua.Yao_5, &bian_gua.Yao_6)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		bian_map["name"] = bian_gua.Name
		bian_map["code"] = bian_gua.Code
		bian_map["alias"] = bian_gua.Alias
		bian_map["location"] = bian_gua.Location
		bian_map["form"] = bian_gua.Form
		bian_map["1"] = bian_gua.Yao_1
		bian_map["2"] = bian_gua.Yao_2
		bian_map["3"] = bian_gua.Yao_3
		bian_map["4"] = bian_gua.Yao_4
		bian_map["5"] = bian_gua.Yao_5
		bian_map["6"] = bian_gua.Yao_6
	}

	//取本卦的爻词解释
	benyaosEx := yao368.Ben_Yaos_Ex{}
	// benyaoRows *Rows
	benyaoRows, err := db.Query("select gua_id, yao_pos, yao_trans from yao386 where gua_id = " + ben_gua.Id);
	if err != nil {
		fmt.Printf("select fail [%s]", err)
	}
	ben_yaos_ex_map := make(map[int]map[string]string)
	for benyaoRows.Next() {
		yao_map := make(map[string]string)
		err = benyaoRows.Scan(&benyaosEx.GuaId, &benyaosEx.YaoPos, &benyaosEx.YaoTrans)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		yao_map["yao_trans"] = benyaosEx.YaoTrans
		ben_yaos_ex_map[benyaosEx.YaoPos]= yao_map
	}

	//取变爻的爻词解释
	bianyaosEx := yao368.Bian_Yaos_Ex{}
	// bianyaoRows *Rows
	bianyaoRows, err := db.Query("select gua_id, yao_pos, yao_trans from yao386 where gua_id = " + bian_gua.Id);
	if err != nil {
		fmt.Printf("select fail [%s]", err)
	}
	bian_yaos_ex_map := make(map[int]map[string]string)
	for bianyaoRows.Next() {
		yao_map := make(map[string]string)
		err = bianyaoRows.Scan(&bianyaosEx.GuaId, &bianyaosEx.YaoPos, &bianyaosEx.YaoTrans)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		yao_map["yao_trans"] = bianyaosEx.YaoTrans
		bian_yaos_ex_map[bianyaosEx.YaoPos]= yao_map
	}

	//根据变爻的数目，确定使用哪个爻辞判定凶吉
	var panCi string
	var panCiExplian string
	switch len(ben_gua.BianYaos) {
	case 0: //本卦卦辞
		panCi = ben_gua.Guaci
		panCiExplian = ben_gua.GuaciFanyi
	case 1: //本卦变爻
		panCi = ben_map[strconv.Itoa(ben_gua.BianYaos[0])]
		panCiExplian = ben_yaos_ex_map[ben_gua.BianYaos[0]]["yao_trans"]
	case 2: //如果卦里有两个爻发生变动，那就用本卦里这两个变爻的占辞来判断吉凶，并以位置靠上的那一个变爻的占辞为主
		panCi = ben_map[strconv.Itoa(ben_gua.BianYaos[1])]
		panCiExplian = ben_yaos_ex_map[ben_gua.BianYaos[1]]["yao_trans"]
	case 3: //本卦变卦卦辞 ，以本卦卦辞为主
		panCi = ben_gua.Guaci
		panCiExplian = ben_gua.GuaciFanyi
	case 4: //变卦两个不变爻
		for i := 1; i < 7; i++ {
			res := common.IssetInSlice(i, ben_gua.BianYaos)
			if res != 0 {
				panCi = bian_map[strconv.Itoa(res)]
				panCiExplian = bian_yaos_ex_map[res]["yao_trans"]
			}
		}
	case 5: //变卦的一个不变爻
		var n int
		for _, i := range ben_gua.BianYaos {
			n += i
		}
		panCi = bian_map[strconv.Itoa(21-n)]
		panCiExplian = bian_yaos_ex_map[21-n]["yao_trans"]

	case 6: // 变卦卦辞
		panCi = bian_gua.Guaci
		panCiExplian = bian_gua.GuaciFanyi
	}
	p := &Res{}
	p.Code = 200
	bianyaos := ben_gua.BianYaos
	p.Data = map[string]interface{}{
		"ben_gua":        ben_map,
		"bian_gua":       bian_map,
		"bian_yaos":      bianyaos,
		"pan_ci":         panCi,
		"pan_ci_explian": panCiExplian,
		"bengua_yaoci_map" : ben_yaos_ex_map,
		"biangua_yaoci_map" : bian_yaos_ex_map,
	}

	p.Status = "Success"
	data, _ := json.Marshal(p)
	c.Ctx.WriteString(string(data))
}
