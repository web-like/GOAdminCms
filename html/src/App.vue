<template>

  <div id="app">
    <com-header></com-header>
    <div id="page-wrapper">
      <div class="row">
        <div class="col-lg-12">
          <h1 class="page-header">{{ $route.name }}</h1>
        </div>
        <!-- /.col-lg-12 -->
      </div>
      <router-view></router-view>
    <!-- /#page-wrapper -->
    </div>
  </div>
  <!-- /#wrapper -->
</template>

<script>
import Header from '@/components/layout/Header'
export default {
  name: 'Index',
  components: {
    'com-header': Header
  }
}
</script>

<style>
/*#app {*/
  /*font-family: 'Avenir', Helvetica, Arial, sans-serif;*/
  /*-webkit-font-smoothing: antialiased;*/
  /*-moz-osx-font-smoothing: grayscale;*/
  /*text-align: center;*/
  /*color: #2c3e50;*/
/*}*/
</style>
