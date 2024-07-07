<script>
    export let handler = () => {};
    export let color = "hsl(100, 50%, 70%)"
</script>

<button on:click={handler} style:background-color={color}><slot /></button>

<style>
    button {
        font-family: inherit;
        border: none;
        padding: 10px 20px;
        color: white;
        border-radius: 5px;
        font-weight: 900;
        font-size: 1.2em;
        width: 130px;
        height: 100px;
    }
</style>