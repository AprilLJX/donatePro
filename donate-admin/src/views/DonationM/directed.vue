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
      <el-table-column align="center" label="捐赠单 ID" width="95">
        <template slot-scope="scope">
          {{ scope.$index }}
        </template>
      </el-table-column>
      <el-table-column label="用户 ID" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.donator_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="项目 ID" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.project_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="是否匿名" width="95" align="center">
        <template slot-scope="scope">
          {{ scope.row.if_anonymous }}
        </template>
      </el-table-column>
      <el-table-column label="捐赠物资" width="300" align="center">
        <template slot-scope="scope">
          {{ scope.row.donate_materials }}
        </template>
      </el-table-column>
      <el-table-column label="留言" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.message }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="状态" width="100" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="捐赠时间" width="150">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.donate_time }}</span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { getTargetDonationList } from '@/api/table'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        进行中: 'success',
        已完成: 'gray',
        待审核: 'danger'
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
      getTargetDonationList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
