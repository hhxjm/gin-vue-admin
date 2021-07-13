<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="pid字段:"><el-input v-model.number="formData.pid" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="name字段:">
                <el-input v-model="formData.name" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="curlevel字段:"><el-input v-model.number="formData.curlevel" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="exp字段:"><el-input v-model.number="formData.exp" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="physicalstrength字段:"><el-input v-model.number="formData.physicalstrength" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="rolemodelframe字段:"><el-input v-model.number="formData.rolemodelframe" clearable placeholder="请输入"></el-input>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createPlayer,
    updatePlayer,
    findPlayer
} from "@/api/player";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Player",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            pid:0,
            name:"",
            curlevel:0,
            exp:0,
            physicalstrength:0,
            rolemodelframe:0,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createPlayer(this.formData);
          break;
        case "update":
          res = await updatePlayer(this.formData);
          break;
        default:
          res = await createPlayer(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findPlayer({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.replayer
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>