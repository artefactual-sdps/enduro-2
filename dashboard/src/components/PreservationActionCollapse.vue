<script setup lang="ts">
import type { api } from "@/client";
import PackageReviewAlert from "@/components/PackageReviewAlert.vue";
import type { EnduroPackagePreservationTask } from "@/openapi-generator";
import StatusBadge from "@/components/StatusBadge.vue";
import { useAuthStore } from "@/stores/auth";
import { onMounted, ref, toRefs } from "vue";

import IconChevronUp from "~icons/bi/chevron-up";

type Card = EnduroPackagePreservationTask & { open: boolean; more: string };

const cards = ref<Card[]>([]);

const authStore = useAuthStore();

const props = defineProps<{
  action: api.EnduroPackagePreservationAction;
  index: number;
}>();

const { action, index } = toRefs(props);

onMounted(() => {
  if (props.action.tasks) {
    for (const task of props.action.tasks.reverse()) {
      let card = <Card>task;
      if (card.note && card.note.includes("\n")) {
        const lines = card.note.split("\n");
        card.note = lines[0];
        card.more = lines.slice(1).join("\n");
      }
      cards.value.push(card);
    }
  }
});

let expandCounter = ref<number>(0);

const isComplete = (task: EnduroPackagePreservationTask) => {
  return task.status == "done" || task.status == "error";
};

const toggleCard = (card: Card, ev: Event) => {
  card.open = !card.open;
  ev.preventDefault();
};
</script>

<template>
  <div class="accordion-item border-0 mb-2">
    <h4 class="accordion-header" :id="'pa-heading-' + index">
      <button
        v-if="action.tasks"
        class="accordion-button collapsed"
        type="button"
        data-bs-toggle="collapse"
        :data-bs-target="'#pa-body-' + index"
        aria-expanded="false"
        :aria-controls="'pa-body-' + index"
      >
        <div class="d-flex flex-column">
          <div class="h4">
            {{ $filters.getPreservationActionLabel(action.type) }}
            <StatusBadge :status="action.status" />
          </div>
          <div>
            <span v-if="action.completedAt">
              Completed
              {{ $filters.formatDateTime(action.completedAt) }}
              (took
              {{
                $filters.formatDuration(action.startedAt, action.completedAt)
              }})
            </span>
            <span v-else>
              Started {{ $filters.formatDateTime(action.startedAt) }}
            </span>
          </div>
        </div>
      </button>
    </h4>
    <div
      v-if="action.tasks"
      :id="'pa-body-' + index"
      class="accordion-collapse collapse bg-light"
      :aria-labelledby="'pa-heading-' + index"
      data-bs-parent="#preservation-actions"
    >
      <div class="accordion-body d-flex flex-column gap-1">
        <PackageReviewAlert
          v-model:expandCounter="expandCounter"
          v-if="authStore.checkAttributes(['package:review'])"
        />
        <div v-for="(item, index) in cards" :key="action.id" class="mb-2 card">
          <div class="card-body">
            <div class="d-flex flex-row align-start gap-3">
              <div class="fd-flex">
                <span
                  class="fs-6 badge rounded-pill border border-primary text-primary"
                >
                  {{ cards.length - index }}
                </span>
              </div>
              <div
                class="d-flex flex-column flex-grow-1 align-content-stretch min-w-0"
              >
                <div class="d-flex flex-wrap pt-1">
                  <div class="me-auto text-truncate fw-bold">
                    {{ item.name }}
                  </div>
                  <div class="me-3">
                    <span
                      v-if="
                        !isComplete(item) &&
                        $filters.formatDateTime(item.startedAt)
                      "
                    >
                      Started:
                      {{ $filters.formatDateTime(item.startedAt) }}
                    </span>
                    <span
                      v-if="
                        isComplete(item) &&
                        $filters.formatDateTime(item.completedAt)
                      "
                    >
                      Completed:
                      {{ $filters.formatDateTime(item.completedAt) }}
                    </span>
                  </div>
                </div>
                <div class="d-flex flex-row gap-4">
                  <div class="flex-grow-1 line-break">
                    {{ item.note }}
                    <span v-if="item.open" :id="'note-' + index + '-more'"
                      ><br />{{ item.more }}</span
                    >
                  </div>
                </div>
              </div>
              <div class="d-flex flex-column pt-1">
                <div>
                  <StatusBadge :status="item.status" />
                </div>
                <div class="align-self-end">
                  <a
                    v-if="item.more"
                    :href="'#note-' + index + '-more'"
                    @click="toggleCard(item, $event)"
                    aria-expanded="false"
                    role="button"
                  >
                    <IconChevronUp />
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
