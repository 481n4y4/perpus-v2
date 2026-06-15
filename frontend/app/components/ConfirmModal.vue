<template>
  <div class="fixed inset-0 z-50 flex items-center justify-center bg-black/50">
    <div
      class="w-full max-w-md rounded-xl bg-gray-800 p-6 text-white shadow-lg"
    >
      <!-- Title -->
      <h2 class="text-xl font-bold">Confirm Delete</h2>

      <!-- Message -->
      <p class="mt-3 text-gray-300">
        Are you sure you want to delete this book?
      </p>

      <!-- Actions -->
      <div class="mt-6 flex justify-end gap-3">
        <button
        @click="emit('close')"
          class="rounded bg-gray-600 px-4 py-2 font-semibold hover:bg-gray-500"
        >
          Cancel
        </button>

        <button
        @click="emit('confirm')"
          class="rounded bg-red-600 px-4 py-2 font-semibold hover:bg-red-500"
        >
          Delete
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
const emit = defineEmits(['close', 'confirm'])
</script>
