<script>
    export let disabled = false;
    export let big = false;
</script>

<button on:click {disabled} class:big={big}><slot /></button>

<style>
    button {
        padding: 0.5em;
        border-radius: 5px;
        border: none;

        background-color: var(--bg, tomato);
        color: var(--fg, white);
    }

    button:hover {
        filter: brightness(1.2);
    }

    button[disabled] {
        filter: grayscale(1);
        cursor: not-allowed;
    }

    button.big {
        font-size: 1.5em;
        font-family: "Urbanist", sans-serif;
        font-weight: 900;
    }
</style>