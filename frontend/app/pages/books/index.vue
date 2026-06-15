<template>
  <div class="bg-gray-900">
    <div class="relative isolate px-6 pt-14 lg:px-8">
      <!-- Dummy Book List - No Logic, Pure HTML & Tailwind CSS -->
      <div class="flex flex-col mx-auto my-10 max-w-3/4">
        
        <div class="flex justify-center my-10" v-if="pendingBooks">
        <div
          class="animate-spin size-10 border-4 border-white border-t-transparent rounded-full"
        ></div>
      </div>

      <div class="flex justify-center my-10" v-else-if="errorBooks">
        <h1 class="w-200 text-center text-red-600 bg-red-200 py-4 rounded">
          Gagal fetching data {{ errorBooks.message }}
        </h1>
      </div>

      <div v-else class="flex flex-col mx-auto my-10 max-w-3/4">
        <!-- Action Buttons -->
        <div class="flex justify-end gap-2 mb-3">
          <button
            class="bg-green-500 hover:bg-green-600 text-white font-semibold px-4 py-2 rounded"
            @click="refreshBooks"
          >
            Refresh
          </button>

          <NuxtLink
            to="/books/add"
            class="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-4 py-2 rounded"
          >
            Add Book
          </NuxtLink>
        </div>
        </div> 

        <!-- Table Container -->
        <div class="p-2 mx-auto text-white w-full">
          <h2 class="mb-3 text-2xl font-semibold leading-tight">Book List</h2>
          <div class="overflow-x-auto">
            <table class="text-xs md:text-sm w-full rounded-lg">
              <thead class="rounded-t-lg bg-gray-700">
                <tr class="text-right">
                  <th class="p-3 text-center">No</th>
                  <th class="p-3 text-center">Name</th>
                  <th class="p-3 text-center">Author</th>
                  <th class="p-3 text-center">Publisher</th>
                  <th class="p-3 text-center">Publish Date</th>
                  <th class="p-3 text-center">Price</th>
                  <th class="p-3 text-center">Stock</th>
                  <th class="p-3 text-center">Cover</th>
                  <th class="p-3 text-center">Action</th>
                 </tr>
              </thead>
              <tbody>
                <tr v-for="(books, index) in books?.data" :key="books.id" class="text-right border-b border-opacity-20 bg-gray-800 hover:bg-gray-700">
                  <td class="px-3 py-4 text-left"><span>{{ index + 1 }}</span></td>
                  <td class="px-3 py-4 text-left"><span>{{ books.name }}</span></td>
                  <td class="px-3 py-4 text-left"><span>{{ books.author}}</span></td>
                  <td class="px-3 py-4 text-left"><span>{{ books.publisher }}</span></td>
                  <td class="px-3 py-4 text-center"><span>{{ books.publish_date }}</span></td>
                  <td class="px-3 py-4 text-left"><span>{{ books.price }}</span></td>
                  <td class="px-3 py-4 text-center"><span>{{ books.stock }}</span></td>
                  <td class="px-3 py-4 text-center">
                    <a href="#" class="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-3 py-2 rounded inline-block text-xs cursor-default">File</a>
                  </td>
                  <td class="px-3 py-2">
                    <div class="flex flex-row gap-1">
                      <a href="#" class="bg-green-600 hover:bg-green-400 text-white font-semibold px-3 py-2 rounded inline-block text-xs cursor-default">Detail</a>
                      <NuxtLink :to="`/books/${books.id}`" class="bg-blue-600 hover:bg-blue-400 text-white font-semibold px-3 py-2 rounded inline-block text-xs cursor-default">Edit</NuxtLink>
                      <button class="bg-red-600 hover:bg-red-400 text-white font-semibold px-3 py-2 rounded text-xs cursor-default">Delete</button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>

const apiBase = useRuntimeConfig().public.apiBase
const {data: books, pending: pendingBooks, error: errorBooks, refresh: refreshBooks} = await useFetch(`${apiBase}/books`)
</script>