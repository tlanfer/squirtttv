<script>
    import {version, runUpdate} from '../lib/version.js';

    let updating = false;
    const initUpdate = async () => {
        updating = true;
        await runUpdate();
        updating = false;
        window.location.reload();
    };
</script>

<div>
    {#if updating}
        Updating, please wait...
    {:else}
        You are running v{$version.version}.
        {#if $version.isLatest}
            This is the latest version.
        {:else }
            Version v{$version.latest} is available.
            <button on:click={initUpdate}>Update now?</button>
        {/if}.
    {/if}
</div>

<style>
    button {
        background: none;
        border: none;
        font-family: inherit;
        padding: 0;
    }

    button:hover {
        text-decoration: underline;
    }
</style>