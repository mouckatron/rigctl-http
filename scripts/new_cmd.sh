#! /bin/bash

if [ $# -eq 0 ]
then
    echo "Missing filename: new_cmd.sh <command>"
    exit 1
fi

OUTPUTFILE="cmd_$1.go"
CAMELCMD=$(sed -r 's/(^|-)(\w)/\U\2/g' <<<"$1")


echo >> $OUTPUTFILE "package main"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "import ("
echo >> $OUTPUTFILE "	"fmt""
echo >> $OUTPUTFILE "	"net/http""
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "	"github.com/gin-gonic/gin""
echo >> $OUTPUTFILE ")"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "// #############################################################################"
echo >> $OUTPUTFILE "// API"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "func apiGet$CAMELCMD(c *gin.Context) {"
echo >> $OUTPUTFILE "	response := rigConnection.cmdGet$CAMELCMD()"
echo >> $OUTPUTFILE "	c.IndentedJSON(http.StatusOK, response)"
echo >> $OUTPUTFILE "}"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "func apiSet$CAMELCMD(c *gin.Context) {"
echo >> $OUTPUTFILE "	var f Frequency"
echo >> $OUTPUTFILE "	if err := c.BindJSON(&f); err != nil {"
echo >> $OUTPUTFILE "		return"
echo >> $OUTPUTFILE "	}"
echo >> $OUTPUTFILE "	response := rigConnection.cmdSet$CAMELCMD(f)"
echo >> $OUTPUTFILE "	c.IndentedJSON(http.StatusOK, response)"
echo >> $OUTPUTFILE "}"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "// #############################################################################"
echo >> $OUTPUTFILE "// API Request/Response Unmarshal/Marshal types"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "// change this"
echo >> $OUTPUTFILE "//type Frequency struct {"
echo >> $OUTPUTFILE "//	Frequency string \`json:"frequency"\`"
echo >> $OUTPUTFILE "//}"
echo >> $OUTPUTFILE "//"
echo >> $OUTPUTFILE "//func (o *Frequency) simpleParse(s string) {"
echo >> $OUTPUTFILE "//	o.Frequency = s"
echo >> $OUTPUTFILE "//}"
echo >> $OUTPUTFILE "//"
echo >> $OUTPUTFILE "// #############################################################################"
echo >> $OUTPUTFILE "// TelnetRigConnection"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "func (c *TelnetRigConnection) cmdGet$CAMELCMD() CommandResponse {"
echo >> $OUTPUTFILE "	raw := c.command("+f")"
echo >> $OUTPUTFILE "	return telnetParseSimpleResponse(raw, &Frequency{})"
echo >> $OUTPUTFILE "}"
echo >> $OUTPUTFILE ""
echo >> $OUTPUTFILE "func (c *TelnetRigConnection) cmdSet$CAMELCMD(f Frequency) CommandResponse {"
echo >> $OUTPUTFILE "	raw := c.command(fmt.Sprintf("+F %s", f.Frequency))"
echo >> $OUTPUTFILE "	return telnetParseSimpleResponse(raw, &f)"
echo >> $OUTPUTFILE "}"
echo >> $OUTPUTFILE ""
