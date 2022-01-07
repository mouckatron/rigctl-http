#! /bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_DIR=$(dirname $SCRIPT_DIR)
UI_IN=$ROOT_DIR/ui/dist/ui
UI_OUT=$ROOT_DIR/cmd/rigctl-http/ui

rm $UI_OUT/*.go

cat <<EOF > $UI_OUT/main.go
package ui

import (
  b64 "encoding/base64"
  "net/http"

	"github.com/gin-gonic/gin"
)

func RouterPaths(r *gin.Engine) {

  r.GET("/ui", func(c *gin.Context) {
    c.Redirect(http.StatusTemporaryRedirect, "/ui/")
  })

  r.GET("/ui/", func(c *gin.Context){
    decoded, _ := b64.StdEncoding.DecodeString(index_html)
    c.Header("Content-Type", "text/html; charset=utf8")
    c.String(http.StatusOK, string(decoded))
  })

EOF


for f in $(ls $UI_IN/*.js $UI_IN/*.css $UI_IN/*.html)
do
    basef=$(basename $f)
    out=$UI_OUT/$basef.b64.go
    varname=$(echo $basef | tr '.' '_' | tr '-' '_')

    echo "package ui" > $out
    echo "" >> $out
    echo -n "var $varname = \`" >> $out
    base64 $f >> $out
    echo "\`" >> $out

    MIME=$(mimetype -b $UI_IN/$f)

    cat <<EOF >> $UI_OUT/main.go
  r.GET("/ui/$basef", func(c *gin.Context){
    decoded, _ := b64.StdEncoding.DecodeString($varname)
    c.Header("Content-Type", "$MIME; charset=utf8")
    c.String(http.StatusOK, string(decoded))
  })

EOF

done

cat <<EOF >> $UI_OUT/main.go
}
EOF
