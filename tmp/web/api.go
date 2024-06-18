package web

func GetApiJsSub() string {
	return `import config from "@/config"
import http from "@/utils/request"

export default {
 index: {
  url: {{GetUrl .Name "index" }},
  name: "列表",
  get: async function (params) {
   return await http.get(this.url, params)
  }
 },
 get: {
  url: {{GetUrl .Name "" }},
  name: "单条信息",
  get: async function (id) {
   return await http.get(this.url + id)
  }
 },
 save: {
  url: {{GetUrl .Name "save" }},
  name: "添加信息",
  post: async function (params) {
   return http.post(this.url, params)
  },
 },
 edit: {
  url: {{GetUrl .Name "edit/" }},
  name: "编辑信息",
  post: async function (id, params) {
   return http.post(this.url + id, params)
  },
 },
 delete: {
  url: {{GetUrl .Name "delete" }},
  name: "删除信息",
  post: async function (params) {
   return http.post(this.url, params)
  },
 }
}`
}
