package templates

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets63ee2efd8d8466106078b867dca773cca0c49385 = "<!DOCTYPE html>\n<html>\n<head>\n  <meta charset=\"UTF-8\">\n  <!-- import CSS -->\n  <link rel=\"stylesheet\" href=\"https://unpkg.com/element-ui/lib/theme-chalk/index.css\">\n  <style type=\"text/css\">\n      .el-table .warning-row {\n          background: oldlace;\n        }\n      .el-table .success-row {\n         background: #f0f9eb;\n      }\n  </style>\n</head>\n<body>\n  <div id=\"app\">\n  <template>\n    <el-table\n      :data=\"tableData\"\n      style=\"width: 100%\"\n      :row-class-name=\"tableRowClassName\">\n      <el-table-column\n        prop=\"date\"\n        label=\"日期\"\n        width=\"180\">\n      </el-table-column>\n      <el-table-column\n        prop=\"name\"\n        label=\"姓名\"\n        width=\"180\">\n      </el-table-column>\n      <el-table-column\n        prop=\"address\"\n        label=\"地址\">\n      </el-table-column>\n    </el-table>\n  </template>\n  </div>\n</body>\n  <!-- import Vue before Element -->\n  <script src=\"https://unpkg.com/vue/dist/vue.js\"></script>\n  <!-- import JavaScript -->\n  <script src=\"https://unpkg.com/element-ui/lib/index.js\"></script>\n  <script>\n    var Main = {\n        methods: {\n          tableRowClassName({row, rowIndex}) {\n            if (rowIndex === 1) {\n              return 'warning-row';\n            } else if (rowIndex === 3) {\n              return 'success-row';\n            }\n            return '';\n          }\n        },\n        data() {\n          return {\n            tableData: [{\n              date: '2016-05-02',\n              name: '王小虎',\n              address: '上海市普陀区金沙江路 1518 弄',\n            }, {\n              date: '2016-05-04',\n              name: '王小虎',\n              address: '上海市普陀区金沙江路 1518 弄'\n            }, {\n              date: '2016-05-01',\n              name: '王小虎',\n              address: '上海市普陀区金沙江路 1518 弄',\n            }, {\n              date: '2016-05-03',\n              name: '王小虎',\n              address: '上海市普陀区金沙江路 1518 弄'\n            }]\n          }\n        }\n      }\n    var Ctor = Vue.extend(Main)\n    new Ctor().$mount('#app')\n  </script>\n</html>"
var _Assetsed8b0b21cd2e2f201192750e67c051b1e773a13c = ""

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{}, map[string]*assets.File{
	"index.tmpl": &assets.File{
		Path:     "index.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1572058806, 1572058806000000000),
		Data:     []byte(_Assets63ee2efd8d8466106078b867dca773cca0c49385),
	}, "assets.go": &assets.File{
		Path:     "assets.go",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1572063207, 1572063207000000000),
		Data:     []byte(_Assetsed8b0b21cd2e2f201192750e67c051b1e773a13c),
	}}, "")
