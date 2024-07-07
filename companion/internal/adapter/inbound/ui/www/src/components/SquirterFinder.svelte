<script>
    import Modal from "./Modal.svelte";
    import FancyButton from "./FancyButton.svelte";
    import {newHosts,setScanning } from "../lib/discovery";
    import {onDestroy, onMount} from "svelte";

    export let addSquirter = (squirter) => {};
    export let showModal = false;

    let manualInput = "";

    let add = (squirter) => {
        addSquirter(squirter);
        showModal = false;
        manualInput = "";
    };

    $: setScanning(showModal)
</script>

<Modal bind:showModal>
    <div>
        <h3>Auto-discovery</h3>

        {#if $newHosts.length === 0}
            <p>No squirters found. Searching...</p>
        {:else}
            <ul>
                {#each $newHosts as squirter}
                    <li>
                        <span><b>Host:</b> {squirter}</span>
                        <FancyButton
                                --bg="hsl(100, 50%, 50%)"
                                on:click={() => add(squirter)}
                        >
                            Add
                        </FancyButton>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>

    <div>
        <h3>Add by hostname / ip address:</h3>
        <form on:submit|preventDefault={()=>{add(manualInput)}}>
            <input type="text" placeholder="Hostname / IP address" bind:value={manualInput} />
            <FancyButton --bg="hsl(100, 50%, 50%)">Add</FancyButton>
        </form>
    </div>
</Modal>

<style>
    form {
        display: flex;
        justify-content: space-between;
        gap: 1em;
    }
    input {
        padding: 0.5em;
        border-radius: 5px;
        border: 1px solid hsl(0, 0%, 90%);
        box-sizing: border-box;
        flex-grow: 1;
    }

    ul {
        padding: 0;
        list-style: none;
    }

    li {
        padding: 0.5em 0;
        border-bottom: 1px solid hsl(0, 0%, 90%);
        display: flex;
        justify-content: space-between;
    }

    li:last-child {
        border-bottom: none;
    }
</style>