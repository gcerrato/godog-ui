<script>
  import { GetReply } from './../wailsjs/go/main/App'
  let messages = $state([
    { sender: "friend", text: "Hi! Type your question" },
  ]);

  let newMessage = $state("");

  async function sendMessage() {
    if (newMessage.trim()) {
      messages.push({ sender: "me", text: newMessage });
      messages.push({ sender: "friend", text: (await GetReply(newMessage)) });
      newMessage = "";
    }
  }
</script>

<div class="flex flex-col w-screen h-screen bg-gray-100">

  <div class="header bg-green-600 text-white p-2.5 text-lg font-bold">Godog</div>


  <div class="messages flex-1 p-2 overflow-y-auto flex flex-col gap-2">
    {#each messages as { sender, text }}
      <div class="message {sender} p-2 rounded-lg text-black {sender === 'me' ? 'bg-green-200 self-end' : 'bg-white self-start'}">
        {text}
      </div>
    {/each}
  </div>


  <div class="footer flex p-2 gap-2 bg-white">
    <input
      type="text"
      placeholder="Type a message..."
      bind:value={newMessage}
      onkeydown={(e) => e.key === "Enter" && sendMessage()}
      class="flex-1 p-2 border border-gray-300 rounded-full outline-none text-black"
    />
    <button onclick={sendMessage} class="bg-green-600 text-white rounded-full w-9 h-9 text-lg cursor-pointer hover:bg-green-500">
      &#9658;
    </button>
  </div>
</div>
