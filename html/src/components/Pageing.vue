<template>
  <nav aria-label="Page navigation">
    <ul class="pagination">
      <li v-bind:class="{disabled: current === 1}">
        <a href="javascript:void(0);" aria-label="Previous" v-on:click="getPageData(1)">
          <span aria-hidden="true">首页</span>
        </a>
      </li>
      <li v-for="(k, v) in indexs" :key="v">
        <a href="javascript:void(0);" @click="getPageData(k)">{{ k }}</a>
      </li>
      <li v-bind:class="{disabled: current===count}">
        <a href="javascript:void(0);" aria-label="Next" v-on:click="getPageData(count)">
          <span aria-hidden="true">尾页</span>
        </a>
      </li>
    </ul>
  </nav>
</template>

<script>
export default {
  name: 'page',
  props: ['count', 'current'],
  data () {
    return {
      showNum: 10,
      indexs: []
    }
  },
  watch: {
    count: function () {
      return this.buildPage()
    }
  },
  methods: {
    getPageData (page) {
      this.$emit('getData', page)
    },
    buildPage () {
      this.indexs = [1]
      let start = Math.round(this.current - this.showNum / 2)
      let end = Math.round(this.current + this.showNum / 2)
      if (start <= 1) {
        start = 2
        end = start + this.showNum - 1
        if (end >= this.count - 1) {
          end = this.count - 1
        }
      }
      //
      if (end >= this.count - 1) {
        end = this.count - 1
        start = end - this.showNum + 1
        if (start <= 1) {
          start = 2
        }
      }
      //
      if (start !== 2) {
        this.indexs.push('...')
      }
      // concat middle page num
      for (let i = start; i <= end; i++) {
        this.indexs.push(i)
      }
      //
      if (end !== this.count - 1) {
        this.indexs.push('...')
      }
      if (this.count !== 1) {
        this.indexs.push(this.count)
      }
    }
  }
}
</script>
