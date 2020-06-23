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
      <el-table-column label="项目名称" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.pro_name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="发起机构" width="110" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.demander_id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="项目简介" width="250" align="center">
        <template slot-scope="scope">
          {{ scope.row.introduction }}
        </template>
      </el-table-column>
      <el-table-column label="项目地址" width="200" align="center">
        <template slot-scope="scope">
          {{ scope.row.rec_address }}
        </template>
      </el-table-column>
      <el-table-column label="需求物资" width="300" align="center">
        <template slot-scope="scope">
          {{ scope.row.materials }}
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="项目类别" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.category | typeFilter">{{ scope.row.category }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="开始时间" width="150">
        <template slot-scope="scope">
          <i class="el-icon-time"/>
          <span>{{ scope.row.init_time }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="截止时间" width="150">
        <template slot-scope="scope">
          <i class="el-icon-time"/>
          <span>{{ scope.row.end_time }}</span>
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="项目状态" width="110" align="center">
        <template slot-scope="scope">
          <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {getProjectList} from '@/api/table'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        进行中: 'success',
        已完成: 'gray',
        待审核:'warning',
        待发布: 'danger'
      }
      return statusMap[status]
    },
    typeFilter(status) {
      const statusMap = {
        生活: 'success',
        教育: 'gray',
        医疗: 'danger'
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
      getProjectList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>
