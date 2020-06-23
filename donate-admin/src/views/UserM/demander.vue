<template>
  <div class="app-container">
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.$index }}
        </template>
      </el-table-column>
      <el-table-column label="账号" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.account }}</span>
        </template>
      </el-table-column>

      <el-table-column label="法人代表" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="机构名称" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.company }}
        </template>
      </el-table-column>
      <el-table-column label="机构类别" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.com_category }}
        </template>
      </el-table-column>
      <el-table-column label="统一信用代码" width="110" align="center">
        <template slot-scope="scope">
          {{ scope.row.credit_code }}
        </template>
      </el-table-column>
      <el-table-column label="地址" width="210" align="center">
        <template slot-scope="scope" style="text-align: left">
          {{ scope.row.com_address }}
        </template>
      </el-table-column>
      <el-table-column label="机构简介" width="210" align="center">
        <template slot-scope="scope">
          {{ scope.row.com_profile }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="状态" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getDemanderList } from '@/api/table'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        审核通过: 'success',
        待审核: 'gray',
        审核未通过: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getDemanderList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
