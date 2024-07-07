<div class="description">
    <span>
        <slot name="description" />
    </span>
    <div class="buttons">
        <slot name="buttons" />
    </div>
</div>

<style>
.description {
    display: flex;
    gap: 2em;
    align-content: flex-start;
    margin-bottom: 2em;

}

.description .buttons * {
    max-width: 200px;
}


</style>