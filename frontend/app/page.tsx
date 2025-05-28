'use client';

import { useEffect, useState } from 'react';

type TODO = {
  id: string;
  title: string;
  completed: boolean;
};

// TODO Do not use useEffects
export default function TodoPage() {
  const [todos, setTodos] = useState<TODO[]>([]);

  // Todo取得
  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const res = await fetch('http://localhost:8080/todos');
    const data = await res.json();
    setTodos(data);
  };

  const handleToggleCompleted = (todo: TODO) => {
    setTodos((prevTodos) => {
      return prevTodos.map((t) => {
        if (t.id === todo.id) {
          return { ...t, completed: !t.completed };
        }
        return t;
      });
    });
  };

  return (
    <div className="p-6 max-w-2xl mx-auto">
      <h1 className="text-2xl font-semibold text-center text-white mb-6">
        📋 Todo List
      </h1>
      <ul className="space-y-4">
        {todos.map((todo) => (
          <li
            key={todo.id}
            className="flex justify-between items-center p-4 bg-gray-800 border border-gray-700 rounded-xl shadow-sm hover:shadow-md transition"
          >
            <div className="flex items-center">
              {/* チェックボックス */}
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
          </li>
        ))}
      </ul>
    </div>
  );
}
