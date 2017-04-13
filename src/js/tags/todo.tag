<todo>
  <h3>{ opts.title || 'Your Tasks' }</h3>

  <ul>
    <li each={ task, i in tasks }>{ task.title }</li>
  </ul>

  <form onsubmit={ add }>
    <input ref="title" placeholder="Task">
    <button>Add</button>
  </form>

  <style>
    todo h3 {
      border-bottom: 2px solid grey;
    }
  </style>

  <script>
    this.tasks = []

    add(e) {
      e.preventDefault()
      var input = this.refs.title
      this.tasks.push({ title: input.value })
      input.value = ''
      input.focus()
    }
  </script>
</todo>
