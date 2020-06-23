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
      <el-table-column align="center" label="ID" width="150">
        <template slot-scope="scope">
          {{ scope.$index }}
        </template>
      </el-table-column>
      <el-table-column label="账号" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.account }}</span>
        </template>
      </el-table-column>
      <el-table-column label="参与捐赠（次）" width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.donate_times }}
        </template>
      </el-table-column>
<!--      <el-table-column class-name="status-col" label="状态" width="150" align="center">-->
<!--        <template slot-scope="scope">-->
<!--          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>-->
<!--        </template>-->
<!--      </el-table-column>-->
    </el-table>
  </div>
</template>

<script>
import { getDonatorList } from '@/api/table'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
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
      getDonatorList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
