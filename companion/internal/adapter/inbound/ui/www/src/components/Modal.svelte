<script>
    export let showModal = false;
    let dialog;

    $: if (dialog && showModal) dialog.showModal();
    $: if (dialog && !showModal) dialog.close();
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
        bind:this={dialog}
        on:close={() => showModal = false}
>
    <button on:click={() => dialog.close()}><span class="material-symbols-outlined">close_small</span></button>
    <slot></slot>
</dialog>

<style>
    button {
        background: none;
        border: none;
        padding: 0.5em 1em;
        border-radius: 5px;
        position: absolute;
        top: 1em;
        right: 1em;
    }
    dialog {
        width: 40em;
        border: none;
        border-radius: 10px;
        box-shadow: 0 0 15px hsla(0, 0%, 0%, 0.2);
        position: relative;

        overflow-x: hidden;
    }
    dialog::backdrop {
        background: hsla(0, 0%, 0%, 0.4);
    }

    dialog[open] {
        animation: zoom .2s cubic-bezier(0.34, 1.56, 0.64, 1);
    }

    dialog[open]::backdrop {
        animation: fade .2s ease-out;
    }

    @keyframes zoom {
        from {
            transform: scale(0.8);
        }
        to {
            transform: scale(1);
        }
    }
    @keyframes fade {
        from {
            opacity: 0;
        }
        to {
            opacity: 1;
        }
    }
</style>