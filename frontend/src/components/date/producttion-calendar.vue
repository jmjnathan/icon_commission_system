<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import { useCommission } from "../../composables/commission/useCommission";

const { commissions, isLoading, fetchAll } = useCommission();

const calendarDate = ref(new Date());

onMounted(async () => {
  await fetchAll();
});

const calendarYear = computed(() => calendarDate.value.getFullYear());

const calendarMonth = computed(() => calendarDate.value.getMonth());

const monthNames = [
  "Januari",
  "Februari",
  "Maret",
  "April",
  "Mei",
  "Juni",
  "Juli",
  "Agustus",
  "September",
  "Oktober",
  "November",
  "Desember",
];

const calendarTitle = computed(
  () => `${monthNames[calendarMonth.value]} ${calendarYear.value}`
);

function previousMonth() {
  calendarDate.value = new Date(calendarYear.value, calendarMonth.value - 1, 1);
}

function nextMonth() {
  calendarDate.value = new Date(calendarYear.value, calendarMonth.value + 1, 1);
}

function goToToday() {
  calendarDate.value = new Date();
}

const calendarDays = computed(() => {
  const year = calendarYear.value;
  const month = calendarMonth.value;

  const firstDay = new Date(year, month, 1);
  const lastDay = new Date(year, month + 1, 0);

  // Senin = 0, Minggu = 6
  const startOffset = (firstDay.getDay() + 6) % 7;

  const totalDays = lastDay.getDate();

  const days: (Date | null)[] = [];

  for (let i = 0; i < startOffset; i++) {
    days.push(null);
  }

  for (let day = 1; day <= totalDays; day++) {
    days.push(new Date(year, month, day));
  }

  return days;
});

function normalizeDate(date: Date) {
  return new Date(date.getFullYear(), date.getMonth(), date.getDate());
}

function isToday(date: Date) {
  const today = new Date();

  return (
    date.getDate() === today.getDate() &&
    date.getMonth() === today.getMonth() &&
    date.getFullYear() === today.getFullYear()
  );
}

function isSameDate(date1: Date, date2: Date) {
  return (
    date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate()
  );
}

function isCommissionActiveOnDate(commission: any, date: Date) {
  if (!commission.order_date || !commission.deadline) {
    return false;
  }

  const orderDate = normalizeDate(new Date(commission.order_date));

  const deadline = normalizeDate(new Date(commission.deadline));

  const current = normalizeDate(date);

  return current >= orderDate && current <= deadline;
}

function isOrderDate(commission: any, date: Date) {
  if (!commission.order_date) {
    return false;
  }

  return isSameDate(new Date(commission.order_date), date);
}

function isDeadline(commission: any, date: Date) {
  if (!commission.deadline) {
    return false;
  }

  return isSameDate(new Date(commission.deadline), date);
}

function getCommissionsForDate(date: Date) {
  return commissions.value.filter((commission: any) =>
    isCommissionActiveOnDate(commission, date)
  );
}

function getStatusClass(status: string) {
  switch (status) {
    case "pending":
      return "status-pending";

    case "in_progress":
      return "status-in_progress";

    case "completed":
      return "status-completed";

    case "cancelled":
      return "status-cancelled";

    default:
      return "status-default";
  }
}

function getStatusLabel(status: string) {
  switch (status) {
    case "pending":
      return "Pending";

    case "in_progress":
      return "Dikerjakan";

    case "completed":
      return "Selesai";

    case "cancelled":
      return "Dibatalkan";

    default:
      return status;
  }
}
</script>

<template>
  <div class="production-calendar">
    <!-- =========================
          HEADER
     ========================== -->
    <div class="calendar-header">
      <div>
        <h2>{{ calendarTitle }}</h2>

        <div class="calendar-summary">
          <span> {{ commissions.length }} commission </span>

          <span v-if="isLoading"> Memuat data... </span>
        </div>
      </div>

      <div class="calendar-actions">
        <button class="nav-button" type="button" @click="previousMonth">
          ‹
        </button>

        <button class="today-button" type="button" @click="goToToday">
          Hari Ini
        </button>

        <button class="nav-button" type="button" @click="nextMonth">›</button>
      </div>
    </div>

    <!-- =========================
          LEGEND
     ========================== -->
    <div class="calendar-legend">
      <div class="legend-item">
        <span class="legend-dot pending"></span>
        Pending
      </div>

      <div class="legend-item">
        <span class="legend-dot in-progress"></span>
        Dikerjakan
      </div>

      <div class="legend-item">
        <span class="legend-dot completed"></span>
        Selesai
      </div>

      <div class="legend-item">
        <span class="legend-deadline"></span>
        Deadline
      </div>
    </div>

    <!-- =========================
          LOADING
     ========================== -->
    <div v-if="isLoading" class="calendar-loading">Memuat commission...</div>

    <!-- =========================
          WEEK DAYS
     ========================== -->
    <div v-else class="calendar-weekdays">
      <div>Sen</div>
      <div>Sel</div>
      <div>Rab</div>
      <div>Kam</div>
      <div>Jum</div>
      <div>Sab</div>
      <div>Min</div>
    </div>

    <!-- =========================
          CALENDAR GRID
     ========================== -->
    <div v-if="!isLoading" class="calendar-grid">
      <div
        v-for="(date, index) in calendarDays"
        :key="index"
        class="calendar-cell"
        :class="{
          'is-today': date && isToday(date),
          'is-empty': !date,
        }">
        <template v-if="date">
          <!-- Date -->
          <div class="date-header">
            <span class="date-number">
              {{ date.getDate() }}
            </span>

            <span
              v-if="getCommissionsForDate(date).length > 0"
              class="commission-count">
              {{ getCommissionsForDate(date).length }}
            </span>
          </div>

          <!-- Commission -->
          <div class="commission-list">
            <div
              v-for="commission in getCommissionsForDate(date)"
              :key="commission.id"
              class="commission-card"
              :class="getStatusClass(commission.status)">
              <!-- Icon / Subject -->
              <div class="commission-name">
                {{ `${commission.items.length} pesanan`}}
              </div>

              <!-- Customer -->
              <div class="commission-customer">
                {{ commission.client.name }}
              </div>

              <!-- Order -->
              <span
                v-if="isOrderDate(commission, date)"
                class="commission-badge order-badge">
                ORDER
              </span>

              <!-- Deadline -->
              <span
                v-if="isDeadline(commission, date)"
                class="commission-badge deadline-badge">
                DEADLINE
              </span>

              <!-- Status -->
              <span class="commission-status">
                {{ getStatusLabel(commission.status) }}
              </span>
            </div>
          </div>
        </template>
      </div>
    </div>

    <!-- =========================
          EMPTY STATE
     ========================== -->
    <div v-if="!isLoading && commissions.length === 0" class="calendar-empty">
      <div class="empty-title">Belum ada commission</div>

      <div class="empty-description">
        Commission yang memiliki order date dan deadline akan muncul di
        kalender.
      </div>
    </div>
  </div>
</template>

<style scoped>
.production-calendar {
  width: 100%;
}

/* =========================
    HEADER
 ========================= */

.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.calendar-header h2 {
  margin: 0;
  font-size: 1.4rem;
  font-weight: 600;
}

.calendar-summary {
  display: flex;
  gap: 1rem;
  margin-top: 0.35rem;
  color: #777;
  font-size: 0.85rem;
}

.calendar-actions {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.calendar-actions button {
  border: 1px solid #ddd;
  background: white;
  border-radius: 6px;
  cursor: pointer;
  transition: 0.15s ease;
}

.calendar-actions button:hover {
  background: #f5f5f5;
}

.nav-button {
  width: 36px;
  height: 36px;
  font-size: 1.4rem;
}

.today-button {
  padding: 0.45rem 0.9rem;
}

/* =========================
    LEGEND
 ========================= */

.calendar-legend {
  display: flex;
  align-items: center;
  gap: 1.25rem;
  margin-bottom: 1rem;
  font-size: 0.8rem;
  color: #666;
}

.legend-item {
  display: flex;
  align-items: center;
  gap: 0.4rem;
}

.legend-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.legend-dot.pending {
  background: #f59e0b;
}

.legend-dot.in-progress {
  background: #3b82f6;
}

.legend-dot.completed {
  background: #22c55e;
}

.legend-deadline {
  width: 8px;
  height: 8px;
  border-radius: 2px;
  border: 2px solid #ef4444;
}

/* =========================
    LOADING
 ========================= */

.calendar-loading {
  padding: 3rem;
  text-align: center;
  color: #777;
  border: 1px solid #eee;
  border-radius: 8px;
}

/* =========================
    WEEKDAYS
 ========================= */

.calendar-weekdays,
.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
}

.calendar-weekdays > div {
  padding: 0.75rem;
  text-align: center;
  font-weight: 600;
  color: #666;
  font-size: 0.85rem;
}

/* =========================
    CALENDAR CELL
 ========================= */

.calendar-cell {
  min-height: 145px;
  border: 1px solid #eee;
  padding: 0.5rem;
  background: white;
  overflow: hidden;
}

.calendar-cell.is-empty {
  background: #fafafa;
}

.calendar-cell.is-today {
  background: #f5f9ff;
}

/* =========================
    DATE
 ========================= */

.date-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.4rem;
}

.date-number {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 26px;
  height: 26px;
  font-size: 0.85rem;
}

.is-today .date-number {
  background: #2563eb;
  color: white;
  border-radius: 50%;
  font-weight: 600;
}

.commission-count {
  font-size: 0.7rem;
  background: #eee;
  color: #666;
  padding: 0.15rem 0.4rem;
  border-radius: 10px;
}

/* =========================
    COMMISSION LIST
 ========================= */

.commission-list {
  display: flex;
  flex-direction: column;
  gap: 0.3rem;
}

/* =========================
    COMMISSION CARD
 ========================= */

.commission-card {
  position: relative;
  padding: 0.45rem;
  border-radius: 5px;
  border-left: 3px solid;
  cursor: pointer;
  transition: 0.15s ease;
  font-size: 0.72rem;
}

.commission-card:hover {
  transform: translateY(-1px);
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
}

/* =========================
    STATUS COLORS
 ========================= */

.commission-card.status-pending {
  background: #fff8e7;
  border-left-color: #f59e0b;
}

.commission-card.status-in_progress {
  background: #eef5ff;
  border-left-color: #3b82f6;
}

.commission-card.status-completed {
  background: #edfdf3;
  border-left-color: #22c55e;
}

.commission-card.status-cancelled {
  background: #fef2f2;
  border-left-color: #ef4444;
}

.commission-card.status-default {
  background: #f5f5f5;
  border-left-color: #999;
}

/* =========================
    COMMISSION TEXT
 ========================= */

.commission-name {
  font-weight: 600;
  line-height: 1.2;
  color: #333;
}

.commission-customer {
  margin-top: 0.15rem;
  color: #777;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.commission-status {
  display: block;
  margin-top: 0.25rem;
  color: #777;
  font-size: 0.65rem;
}

/* =========================
    BADGES
 ========================= */

.commission-badge {
  display: inline-block;
  margin-top: 0.3rem;
  margin-right: 0.15rem;
  padding: 0.1rem 0.3rem;
  border-radius: 3px;
  font-size: 0.58rem;
  font-weight: 700;
  letter-spacing: 0.03em;
}

.order-badge {
  background: #dbeafe;
  color: #1d4ed8;
}

.deadline-badge {
  background: #fee2e2;
  color: #b91c1c;
}

/* =========================
    EMPTY STATE
 ========================= */

.calendar-empty {
  padding: 3rem;
  text-align: center;
  border: 1px solid #eee;
  border-top: none;
}

.empty-title {
  font-weight: 600;
  color: #444;
}

.empty-description {
  margin-top: 0.35rem;
  font-size: 0.85rem;
  color: #888;
}
</style>
