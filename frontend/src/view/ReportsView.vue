<script setup>
import HeaderComponent from "../components/Header/HeaderComponent.vue";
import { ref, onMounted } from 'vue';
import { GetAllReports } from '../../wailsjs/go/service/Service.js';

const reports = ref([]);
const loading = ref(true);
const error = ref(null);

const fetchReports = async () => {
  try {
    reports.value = await GetAllReports();
  } catch (err) {
    console.error('Error fetching reports:', err);
  } 
};

onMounted(() => {
  fetchReports();
});
</script>

<template>
  <HeaderComponent />
  <div class="container">
    <h1 class="title">Reports</h1>
    <div сlass="reports-grid">
      <div v-for="report in reports" :key="report.id" class="report-card">
        <h3>{{ report.title }}</h3>
        <p class="description">{{ report.description }}</p>
        <div class="report-footer">
          <span class="title" :class="report.title">{{ report.title }}</span>
          <span class="description" :class="report.description">{{ report.description }}</span>
          <span class="status" :class="report.status">{{ report.status }}</span>
          <span class="priority">{{ report.priority }}</span>
          <span class="category">{{ report.category }}</span>
          <span class="tags">{{ report.tags }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  background: #fff;
  border-radius: 15px;
  width: 100%;
  height: 100%;
  padding: 20px 20px;
}

.title {
  font-size: 24px;
  margin-bottom: 20px;
  color: #333;
}

.loading, .no-reports, .error {
  text-align: center;
  padding: 20px;
  color: #666;
}

.error {
  color: #dc3545;
}

.retry-btn {
  margin-top: 10px;
  padding: 8px 16px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.retry-btn:hover {
  background-color: #0056b3;
}

.reports-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
  padding: 10px;
}

.report-card {
  background: #f8f9fa;
  border-radius: 8px;
  padding: 15px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s;
}

.report-card:hover {
  transform: translateY(-2px);
}

.report-card h3 {
  margin: 0 0 10px 0;
  color: #2c3e50;
}

.description {
  color: #666;
  margin-bottom: 15px;
  font-size: 0.9em;
}

.report-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 10px;
}

.status {
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 0.8em;
  font-weight: 500;
}

.status.pending {
  background: #fff3cd;
  color: #856404;
}

.status.completed {
  background: #d4edda;
  color: #155724;
}

.status.inprogress {
  background: #cce5ff;
  color: #004085;
}

.date {
  font-size: 0.8em;
  color: #6c757d;
}
</style>