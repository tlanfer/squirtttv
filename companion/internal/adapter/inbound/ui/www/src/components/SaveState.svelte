<script>
    import {dirty} from "../lib/config.js";
</script>

{#if $dirty}
    <div class="status" class:dirty={$dirty}>Unsaved changes</div>
{:else}
    <div class="status">Config saved</div>
{/if}

<style>
    .status {
        margin-top: 5px;
    }

    .status:before {
        content: "● ";
        color: green;
    }

    .status.dirty:before {
        color: red;
    }
</style>