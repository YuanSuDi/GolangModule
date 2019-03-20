package vcf


type Contact struct {
	name string
	Header string //头部标签
	Footer string //脚部标签
	cellPhone string //手机号
	telePhone1 string //电话1
	telePhone2 string //电话2
}

//读取一行
func (c *Contact) ReadRow(row []string){
	for i := 0; i< len(row); i++ {
		switch i {
		case 0:
			c.SetName(row[i])
		case 1:
			c.SetCellPhone(row[i])
		}
	}
}


func (c *Contact) SetName(n string)  {
	c.name = "FN:" + n
}

func (c *Contact) GetName() string {
	return c.name
}

func (c *Contact) TelePhone1() string {
	return c.telePhone1
}

func (c *Contact) SetTelePhone1(telePhone1 string) {
	c.telePhone1 = telePhone1
}


func (c *Contact) GetCellPhone() string {
	return c.cellPhone
}

func (c *Contact) SetCellPhone(p string){
	c.cellPhone = "TEL;TYPE=CELL:" + p
}

