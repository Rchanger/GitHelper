<template>
  <div class="hello">
        Commit Message : <input v-model="message" placeholder="message">
        File : <input v-model="file" placeholder="file">
        <button @click=createCommit> Commit </button>
  </div>
</template>

<script>
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios, axios)
export default {
  name: 'CommitForm',
  data () {
    return {
        message: "",
        file: "",
        repo:"",
        branch:""
    }
  },
  mounted() {
      this.repo = this.$route.query.repo
      this.branch = this.$route.query.branch
  },
  methods: {
      createCommit() {
        var files = []
        files.push(this.file)
        var request = {"repo":this.repo, "branch": this.branch,"commit_message":this.message,"files":files}
        Vue.axios.post("http://localhost:8080/commit", request).then((response) => {
          alert("Successful commit")
          this.message = ""
          this.file = ""
          this.$router.push({name:'Branch', params:{repo:this.repo}})
        })
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>