<script setup lang="ts">


</script>

<template>
    <nav>
        <RouterLink to="/" class="logo">
            <img src="../../assets/sb.png" alt="logo">
            <div class="logo-item">
                Scoreboard
            </div>
        </RouterLink>
        <div class="links">
            <RouterLink to="/" class="link-item">Home</RouterLink>
            <RouterLink to="/register" class="link-item">Register</RouterLink>
            <RouterLink to="/generate-questions" class="link-item">Generate Questions</RouterLink>
        </div>
    </nav>
</template>

<style scoped lang="scss">
 nav{
    display: flex;
    justify-content: space-between;
    background-color: var(--main-text-color-dark-transparent);
    border: 3px solid var(--primary-color);

    position: sticky;
    top: 0;
    z-index: 1000; 
  }

  .logo{
    display: flex;

    // border: 2px solid white;
    width: fit-content;

    justify-content: center;
    align-items: center;
    margin: 15px;
    font-size: 25px;
    text-decoration: none;

    img{
      height: 30px;
      width: auto;
      margin-right: 8px;
    }

    .logo-item{
      color: var(--main-bg-color);
      transition: 0.5s;

      &:hover{
        color: var(--primary-color);
      }
    }

    &:hover{
      cursor: pointer;
    }
  }

  .links{
    display: flex;
    justify-content: center;
    align-items: center;
    width: fit-content;
    // border: 2px red solid;

    .link-item{
      display: flex;
      align-items: center;

      font-size: 18px;
      // margin: 15px;
      padding: 0px 20px;

      color: var(--primary-color);
      // border: 2px solid black;

      transition: 0.5s;
      height: 100%;

      &:hover{
        background-color: blue;
      }
    }
  }
</style>