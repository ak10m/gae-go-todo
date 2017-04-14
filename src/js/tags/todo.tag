<todo>
  <h3>{ opts.title || 'Your Tasks' }</h3>

  <ul>
    <li each={ task, i in tasks }>
      <a href="/todo" onclick={ deleteTask }>x</a>
      { task.title }
    </li>
  </ul>

  <form onsubmit={ addTask } ref="newForm">
    <input name="title" placeholder="Task">
    <button>Add</button>
  </form>

  <style>
    todo h3 {
      border-bottom: 2px solid grey;
    }
  </style>

  <script>
    self = this

    loadTasks() {
      fetch('/todo')
      .then( response => { return response.json() })
      .then( json => {
        self.tasks = json
        self.update()
      })
    }

    addTask(e) {
      e.preventDefault()

      let newForm = this.refs.newForm
      if (newForm.title) {
        fetch('/todo', {
          method: 'POST',
          body: new FormData(newForm)
        })
        .then( response => {
          return response.json()
        })
        .then( json => {
          self.tasks.push(json)
          self.update()
        })
      }
      newForm.title.value = ''
      newForm.title.focus()
    }

    deleteTask(e) {
      e.preventDefault()

      let link = e.target
      let url = link.href

      fetch(url, { method: 'DELETE' })
      .then( response => {
        if (response.ok) {
          let item = link.parentNode
          let list = item.parentNode
          list.removeChild(item)
        }
      })
    }

    this.on('mount', () => {
      self.loadTasks()
    })
  </script>
</todo>
