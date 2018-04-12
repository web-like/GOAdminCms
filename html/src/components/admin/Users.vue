<template>
  <!-- /.row -->
  <div class="row" id="users">
    <div class="col-lg-12">
      <div class="panel panel-default">
        <!-- /.panel-heading -->
        <div class="panel-body">
          <div class="row" style="margin-bottom: 15px">
            <div class="col-lg-9">
              <button type="button" class="btn btn-success btn-sm" v-on:click="showModal('add')">添加</button>
            </div>
            <div class="col-lg-3">
              <div class="input-group">
                <input type="text" class="form-control" placeholder="请输入用户组..." v-model="searchData.Title">
                <span class="input-group-btn">
                  <button class="btn btn-primary" type="button" @click="lists(1)">检索</button>
                </span>
              </div><!-- /input-group -->
            </div>
          </div>
          <div class="table-responsive">
            <table class="table table-striped table-bordered table-hover">
              <thead>
              <tr>
                <th>ID</th>
                <th>用户名</th>
                <th>邮箱</th>
                <th>用户组</th>
                <th>创建人</th>
                <th>创建时间</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="(item, index) in items"  :key="item.Id">
                <td>{{ item.id }}</td>
                <td>{{ item.username }}</td>
                <td>{{ item.email }}</td>
                <td>{{ item.role_id }}</td>
                <td>{{ item.parent_id }}</td>
                <td>{{ item.created_at }}</td>
                <td>
                  <button type="button" class="btn btn-info btn-xs" v-on:click="showModal(index)">编辑</button>
                  <button type="button" class="btn btn-danger btn-xs" v-on:click="del(item.Id)">删除</button>
                </td>
              </tr>
              </tbody>
            </table>

            <!-- 分页 -->
            <my-page :count="page.total" :current="page.current" v-on:getData="lists"></my-page>
            <!-- /分页 -->

          </div>
          <!-- /.table-responsive -->
        </div>
        <!-- /.panel-body -->
      </div>
    </div>
    <!-- /.row -->
    <!-- Modal -->
    <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myModalLabel">{{ modalTitle }}</h4>
          </div>
          <div class="modal-body form-horizontal">
            <!-- -->
            <div class="form-group">
              <label for="username" class="col-sm-3 control-label">用户名：</label>
              <div class="col-sm-8">
                <input type="text" v-model="post_data.username" class="form-control" id="username">
              </div>
              <div v-show="this.err_info.title !== undefined" class="col-sm-offset-3 col-sm-8 text-danger">
                {{ this.err_info.title }}
              </div>
            </div>
            <!-- -->
            <div class="form-group">
              <label for="email" class="col-sm-3 control-label">邮箱：</label>
              <div class="col-sm-8">
                <input type="email" v-model="post_data.emial" class="form-control" id="email">
              </div>
              <div v-show="this.err_info.title !== undefined" class="col-sm-offset-3 col-sm-8 text-danger">
                {{ this.err_info.title }}
              </div>
            </div>
            <!-- -->
            <div class="form-group">
              <label for="role" class="col-sm-3 control-label">分组：</label>
              <div class="col-sm-8">
                <select id="role" class="form-control" v-model="post_data.role_id">
                  <option v-for="(role, role_index) in roles" value="1" :key="role_index">{{ role.Title }}</option>
                </select>
              </div>
            </div>
            <!-- -->
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            <button type="button" class="btn btn-primary" v-on:click="addOrSave">保存</button>
          </div>
        </div>
      </div>
    </div>
    <!-- 删除框 -->
    <div class="modal fade bs-example-modal-sm" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel">
      <div class="modal-dialog modal-sm" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title">执行。。。</h4>
          </div>
          <div class="modal-body">
            {{ del_msg }}
          </div>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
import Page from '@/components/Pageing'
export default {
  name: 'users',
  components: {
    'my-page': Page
  },
  data () {
    return {
      httpUrl: 'user',
      items: '',
      // 弹出框标题
      modalTitle: '',
      // 添加编辑数据源
      post_data: {
        id: 0,
        username: '',
        email: '',
        role_id: 1
      },
      err_info: '', // 错误信息
      del_msg: '', // 删除提示信息
      page: { // 分页信息
        current: 1,
        limit: 14,
        offset: 0,
        total: 0
      },
      // 检索
      searchData: {
        Title: ''
      },
      // role_info
      roles: ''
    }
  },
  created: function () {
    this.lists(this.page.current)
  },
  mounted: function () {
    let roleUrl = 'role'
    let parms = {limit: 1000, offset: 0}
    this.$http.get(roleUrl + this.buildHttp(parms)).then(res => {
      this.roles = res.data
    })
  },
  methods: {
    lists (page) { // 列表
      if (page === undefined) {
        page = 1
      }
      this.page.current = page
      this.page.offset = (this.page.current - 1) * this.page.limit
      let httpQuery = {'offset': this.page.offset, 'limit': this.page.limit}
      httpQuery = {...httpQuery, ...this.searchData}
      this.$http.get(this.httpUrl + this.buildHttp(httpQuery)).then(res => {
        this.items = res.data
        this.page.total = Math.ceil(res.count / this.page.limit)
      })
    },
    // 添加或者保存
    addOrSave () {
      this.$http.post(this.httpUrl, this.post_data).then(res => {
        if (res.err_field !== undefined) {
          this.err_info = res.err_field
        } else {
          this.closeModal()
          if (this.post_data.Id === 0) {
            this.lists()
          }
        }
      }).catch(function (err) {
        console.log(err)
      })
    },
    // 删除记录
    del (id) {
      this.$http.delete(this.httpUrl + '/' + id).then(res => {
        $('.bs-example-modal-sm').modal('show')
        this.del_msg = res.msg
        this.lists()
        window.setTimeout(function () {
          $('.bs-example-modal-sm').modal('hide')
        }, 1000)
      }).catch(function (err) {
        console.log(err)
      })
    },
    showModal (index) {
      if (index !== 'add') {
        this.modalTitle = '编辑'
        this.post_data = this.items[index]
      } else {
        this.modalTitle = '添加'
        this.post_data = {
          Id: 0,
          Title: ''
        }
      }
      this.err_info = ''
      $('#myModal').modal('show')
    },
    closeModal () {
      $('#myModal').modal('hide')
    },
    buildHttp (jsonStr) {
      let str = '?'
      for (var i in jsonStr) {
        if (str !== '?') {
          str += '&'
        }
        str += i + '=' + jsonStr[i]
      }
      return str
    }
  }
}
</script>
