<template>
  <div class="flex flex-col justify-between bg-gray-900 min-h-screen">
    <Navbar item1="Home" item2="About" item3="Book Lists" />
    <slot />
    <Footer />
  </div>
</template>
