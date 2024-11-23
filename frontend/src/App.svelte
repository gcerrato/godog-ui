<script lang="ts">
  import { GetReply, SetDirectory, UploadDocuments } from './../wailsjs/go/main/App';
  import snarkdown from 'snarkdown';
  import { tick } from 'svelte';

  let messages = $state([
    { sender: "friend", text: "Hi! Type your question" },
  ]);

  let newMessage = $state("");
  let loading = $state(false);
  let currentDir = $state("./");

  let messagesContainer: HTMLElement;

  async function sendMessage() {
    if (newMessage.trim()) {
      messages.push({ sender: "me", text: newMessage });
      loading = true;
      scrollToBottom();
      let prompt = newMessage;
      newMessage = "";
      messages.push({ sender: "friend", text: (await GetReply(prompt)) });
      loading = false;
      await tick();
      scrollToBottom();
    }
  }

  function scrollToBottom() {
    if (messagesContainer) {
      messagesContainer.scrollTop = messagesContainer.scrollHeight;
    }
  }

  function init(el: HTMLElement){
    el.focus()
  }

  async function setDirectory() {
    currentDir = await SetDirectory()
  }

  async function processDirectory() {
    UploadDocuments()
  }

</script>

<div class="flex flex-col w-screen h-screen bg-gray-100 dark:bg-gray-900">

  <div class="header bg-green-600 text-white p-2.5 text-lg font-bold dark:bg-green-700 text-left">Current directory: {currentDir}</div>

  <div bind:this={messagesContainer} class="messages flex-1 p-2 overflow-y-auto flex flex-col gap-2">
    {#each messages as { sender, text }}
      <div
        class="message {sender} p-2 rounded-lg
          {sender === 'me' ? 'bg-green-200 text-black self-end dark:bg-green-700 dark:text-white' :
          'bg-white text-black self-start dark:bg-gray-800 dark:text-gray-200'} text-left">
        {#if sender === 'me'}
          {text}
        {:else}
          {@html snarkdown(text)}
        {/if}
      </div>
    {/each}
    {#if loading}
      <div class="loading p-2 text-gray-500 dark:text-gray-400">Loading...</div>
    {/if}
  </div>

  <div class="footer flex p-2 gap-2 bg-white dark:bg-gray-800">
    <button
      aria-label="set the directory"
      onclick="{setDirectory}"
      title="set the directory"
      class="rounded-full w-9 h-9 text-lg cursor-pointer"
    >
      <svg class="w-6 h-6 text-gray-800 dark:text-green-600 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512"> <path d="M384 480l48 0c11.4 0 21.9-6 27.6-15.9l112-192c5.8-9.9 5.8-22.1 .1-32.1S555.5 224 544 224l-400 0c-11.4 0-21.9 6-27.6 15.9L48 357.1 48 96c0-8.8 7.2-16 16-16l117.5 0c4.2 0 8.3 1.7 11.3 4.7l26.5 26.5c21 21 49.5 32.8 79.2 32.8L416 144c8.8 0 16 7.2 16 16l0 32 48 0 0-32c0-35.3-28.7-64-64-64L298.5 96c-17 0-33.3-6.7-45.3-18.7L226.7 50.7c-12-12-28.3-18.7-45.3-18.7L64 32C28.7 32 0 60.7 0 96L0 416c0 35.3 28.7 64 64 64l23.7 0L384 480z"/> </svg>
    </button>
    <button
      aria-label="process directory"
      onclick="{processDirectory}"
      title="Process the directory"
      class="rounded-full w-9 h-9 text-lg cursor-pointer"
    >
      <svg class="w-6 h-6 text-gray-800 dark:text-green-600 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--!Font Awesome Free 6.7.1 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2024 Fonticons, Inc.--><path d="M176 24c0-13.3-10.7-24-24-24s-24 10.7-24 24l0 40c-35.3 0-64 28.7-64 64l-40 0c-13.3 0-24 10.7-24 24s10.7 24 24 24l40 0 0 56-40 0c-13.3 0-24 10.7-24 24s10.7 24 24 24l40 0 0 56-40 0c-13.3 0-24 10.7-24 24s10.7 24 24 24l40 0c0 35.3 28.7 64 64 64l0 40c0 13.3 10.7 24 24 24s24-10.7 24-24l0-40 56 0 0 40c0 13.3 10.7 24 24 24s24-10.7 24-24l0-40 56 0 0 40c0 13.3 10.7 24 24 24s24-10.7 24-24l0-40c35.3 0 64-28.7 64-64l40 0c13.3 0 24-10.7 24-24s-10.7-24-24-24l-40 0 0-56 40 0c13.3 0 24-10.7 24-24s-10.7-24-24-24l-40 0 0-56 40 0c13.3 0 24-10.7 24-24s-10.7-24-24-24l-40 0c0-35.3-28.7-64-64-64l0-40c0-13.3-10.7-24-24-24s-24 10.7-24 24l0 40-56 0 0-40c0-13.3-10.7-24-24-24s-24 10.7-24 24l0 40-56 0 0-40zM160 128l192 0c17.7 0 32 14.3 32 32l0 192c0 17.7-14.3 32-32 32l-192 0c-17.7 0-32-14.3-32-32l0-192c0-17.7 14.3-32 32-32zm192 32l-192 0 0 192 192 0 0-192z"/></svg>
    </button>
    <input
      type="text"
      placeholder="Type a message..."
      bind:value={newMessage}
      use:init
      onkeydown={(e) => e.key === "Enter" && sendMessage()}
      class="flex-1 p-2 border border-gray-300 rounded-full outline-none text-black dark:bg-gray-700 dark:border-gray-600 dark:text-white"
    />
    <button
      aria-label="send message"
      onclick={sendMessage}
      class="w-9 h-9 text-lg cursor-pointer">
      <svg class="w-6 h-6 text-gray-800 dark:text-green-600 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" >       <path d="M0 256a256 256 0 1 1 512 0A256 256 0 1 1 0 256zM188.3 147.1c-7.6 4.2-12.3 12.3-12.3 20.9l0 176c0 8.7 4.7 16.7 12.3 20.9s16.8 4.1 24.3-.5l144-88c7.1-4.4 11.5-12.1 11.5-20.5s-4.4-16.1-11.5-20.5l-144-88c-7.4-4.5-16.7-4.7-24.3-.5z"/> </svg>
    </button>

  </div>
</div>
