package zones

func (c *control) Update(ip string, domainName string) (int64, error) {
	var sqlCommand = "UPDATE pdns.records SET content = ? WHERE name = ?"
	var par = []interface{}{ip, domainName}
	id, err := c.mysqlP.UpdateOrDelete(sqlCommand, par...)
	if err != nil {
		return -1, err
	}
	return id, nil
}
