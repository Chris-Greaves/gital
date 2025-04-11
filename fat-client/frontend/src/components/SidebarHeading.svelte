<script lang="ts">
    let { children, name } = $props();
</script>


<div class="">
    <h2 id="heading" class="border-b-2 border-b-cyan-900 my-1">{name}</h2>
    <ul class="space-y-2">
        {@render children?.()}
    </ul>
</div>

<style>
    #heading {
        font-size: large;
    }
</style>