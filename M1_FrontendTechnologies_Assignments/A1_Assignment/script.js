const taskInput = document.getElementById('taskInput');
const addTaskBtn = document.getElementById('addTaskBtn');
const taskList = document.getElementById('taskList');
const pendingTasks = document.getElementById('pendingTasks');
let tasks = JSON.parse(localStorage.getItem('tasks')) || [];
function renderTasks() {
  taskList.innerHTML = '';
  tasks.forEach((task, index) => {
    const taskItem = document.createElement('li');
    taskItem.className = `list-group-item d-flex justify-content-between align-items-center ${task.completed ? 'completed' : ''}`;
    taskItem.innerHTML = `
      <span>${task.name}</span>
      <div>
        <button 
          class="btn btn-sm ${task.completed ? 'btn-success' : 'btn-danger'} me-1" 
          onclick="toggleComplete(${index})">
          <i class="bi bi-check-lg"></i>
        </button>
        <button class="btn btn-sm btn-primary me-1" onclick="editTask(${index})">
          <i class="bi bi-pencil"></i>
        </button>
        <button class="btn btn-sm btn-danger" onclick="deleteTask(${index})">
          <i class="bi bi-trash"></i>
        </button>
      </div>
    `;
    taskList.appendChild(taskItem);
  });
  updatePendingCount();
  localStorage.setItem('tasks', JSON.stringify(tasks));
}
addTaskBtn.addEventListener('click', () => {
  const taskName = taskInput.value.trim();
  if (taskName) {
    tasks.push({ name: taskName, completed: false });
    taskInput.value = '';
    renderTasks();
  }
});
function toggleComplete(index) {
  tasks[index].completed = !tasks[index].completed;
  renderTasks();
}
function editTask(index) {
  const newTaskName = prompt('Edit your task:', tasks[index].name);
  if (newTaskName !== null && newTaskName.trim() !== '') {
    tasks[index].name = newTaskName.trim();
    renderTasks();
  }
}
function deleteTask(index) {
  tasks.splice(index, 1);
  renderTasks();
}
function updatePendingCount() {
  const pendingCount = tasks.filter(task => !task.completed).length;
  pendingTasks.textContent = `Pending tasks: ${pendingCount}`;
}
renderTasks();
