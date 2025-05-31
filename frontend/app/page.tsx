'use client';

import { useEffect, useState } from 'react';

type TODO = {
  id: string;
  title: string;
  completed: boolean;
};

export default function TodoPage() {
  const [todos, setTodos] = useState<TODO[]>([]);
  const [input, setInput] = useState<string>('');

  // Todo取得
  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const res = await fetch('http://localhost:8080/todos');
    const data: Array<TODO> = await res.json();
    // 新しいTODOが上に来るように
    data.sort((x: TODO, y: TODO) => parseInt(y.id) - parseInt(x.id));
    setTodos(data);
  };

  const handleToggleCompleted = async (todo: TODO) => {
    todo.completed = !todo.completed;
    await fetch('http://localhost:8080/todos', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(todo),
    });

    //入れた値をリセット
    setInput('');
    //全Todoリストを再取得する
    fetchTodos();
  };

  const handleAddTodo = async () => {
    await fetch('http://localhost:8080/todos', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        title: input,
        completed: false,
      }),
    });

    //入れた値をリセット
    setInput('');
    //全Todoリストを再取得する
    fetchTodos();
  };

  const handleDeleteTodo = async (todo: TODO) => {
    console.log(todo);
    await fetch(`http://localhost:8080/todos/${todo.id}`, {
      method: 'DELETE',
    });

    fetchTodos();
  };

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <h1 className="text-2xl font-semibold text-center text-white mb-6">
        📋 Todo List
      </h1>
      <div className="flex items-center gap-2 mb-6">
        <input
          className="flex-1 p-2 rounded bg-gray-700 text-white placeholder-gray-400"
          type="text"
          placeholder="新しいTodoを追加..."
          value={input}
          onChange={(e) => setInput(e.target.value)}
        />
        <button
          type="button"
          onClick={handleAddTodo}
          className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
        >
          追加
        </button>
      </div>

      <ul className="space-y-4">
        {todos.map((todo) => (
          <li
            key={todo.id}
            className="flex justify-between items-center p-4 bg-gray-800 border border-gray-700 rounded-xl shadow-sm hover:shadow-md transition"
          >
            <div className="flex items-center">
              <input
                type="checkbox"
                checked={todo.completed}
                onChange={() => handleToggleCompleted(todo)}
                className="w-5 h-5 mr-4 accent-green-600"
              />
              <span className="text-lg font-medium text-white">
                {todo.title}
              </span>
            </div>
            <button
              onClick={() => handleDeleteTodo(todo)}
              className={`px-3 py-1 rounded-full text-sm font-semibold text-white`}
            >
              削除
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
