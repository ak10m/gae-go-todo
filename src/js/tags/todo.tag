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
      var titleInput = this.refs.title
      var title = titleInput.value
      if (title) {
        this.tasks.push({ title: title })
      }
      titleInput.value = ''
      titleInput.focus()
    }
  </script>
</todo>
