<script>
    import {version, runUpdate} from '../lib/version.js';

    let state = $version.isLatest ? "ok" : "old";
    const initUpdate = async () => {
        state = "updating";
        const success = await runUpdate();
        if( success ) {
            state = "updated";
        } else {
            state = "error";
        }
    };
</script>

<div>
    You are running {$version.version}.
    {#if state==="ok"}
        This is the latest version.
    {/if}
    {#if state==="old"}
        Version {$version.latest} is available.
        <button on:click={initUpdate}>Update now?</button>
    {/if}
    {#if state==="updating"}
        Updating, please wait...
    {/if}
    {#if state==="updated"}
        Updated to v{$version.latest}. Please restart the companion
    {/if}
    {#if state==="error"}
        An error occurred. Please try again later.
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