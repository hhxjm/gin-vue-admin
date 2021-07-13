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
           <el-button v-if="this.wf.clazz == 'start'" @click="start" type="primary">启动</el-button>
           // complete传入流转参数 决定下一步会流转到什么位置 此处可以设置多个按钮来做不同的流转
           <el-button v-if="this.wf.clazz == 'userTask'" @click="complete('yes')" type="primary">提交</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    startWorkflow,
    completeWorkflowMove
} from "@/api/workflowProcess";
import infoList from "@/mixins/infoList";
import { mapGetters } from "vuex";
export default {
  name: "Player",
  mixins: [infoList],
  props:{
      business:{
         type:Object,
        default:function(){return null}
      },
      wf:{
        type:Object,
        default:function(){return{}}
      },
      workflowMoveID:{
        type:Number,
        default:0
      }
   },
  data() {
    return {formData: {
            pid:0,
            name:"",
            curlevel:0,
            exp:0,
            physicalstrength:0,
            rolemodelframe:0,
            
      }
    };
  },
  computed:{
      canShow(){
        if(this.wf.assignType == "user"){
          if(this.wf.assginValue.indexOf(","+this.userInfo.ID+",")>0){
            return true
          }else{
            return false
          }
        }else if(this.wf.assign_type == "authority"){
          if(this.wf.assginValue.indexOf(","+this.userInfo.authorityId+",")>0){
            return true
          }else{
            return false
          }
        }
      },
      ...mapGetters("user", ["userInfo"])
  },
  methods: {
    async start() {
      const res = await startWorkflow({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessId:0,
              businessType:"player",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"create",
              param:""
              }
          });
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"启动成功"
        })
       this.back()
      }
    },
    async complete(param){
     const res = await completeWorkflowMove({
            business:this.formData,
            wf:{
              workflowMoveID:this.workflowMoveID,
              businessID:this.formData.ID,
              businessType:"player",
              workflowProcessID:this.wf.workflowProcessID,
              workflowNodeID:this.wf.id,
              promoterID:this.userInfo.ID,
              operatorID:this.userInfo.ID,
              action:"complete",
              param:param
              }
     })
     if(res.code == 0){
       this.$message({
          type:"success",
          message:"提交成功"
       })
       this.back()
     }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
    if(this.business){
     this.formData = this.business
    }
}
};
</script>

<style>
</style>