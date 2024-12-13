document.getElementById("addTaskButton").addEventListener("click", function () {
  const taskInput = document.getElementById("taskInput");
  const task = taskInput.value;

  if (task) {
    fetch("/tasks", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ task: task }),
    })
      .then((response) => response.json())
      .then((data) => {
        addTaskToList(data.task);
        taskInput.value = "";
      });
  }
});

function addTaskToList(task) {
  const taskList = document.getElementById("taskList");
  const li = document.createElement("li");
  li.textContent = task;

  const deleteButton = document.createElement("button");
  deleteButton.textContent = "Удалить";
  deleteButton.className = "delete";
  deleteButton.onclick = function () {
    deleteTask(task, li);
  };

  li.appendChild(deleteButton);
  taskList.appendChild(li);
}

function deleteTask(task, li) {
  fetch("/tasks", {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ task: task }),
  }).then((response) => {
    if (response.ok) {
      li.remove();
    }
  });
}

const socket = new WebSocket(`ws://${window.location.host}/ws`);

socket.onmessage = function (event) {
  if (event.data === "reload") {
    location.reload(); // Перезагрузка страницы
  }
};
