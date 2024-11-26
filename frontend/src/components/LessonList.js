import React, { useState, useEffect } from 'react';
import { getLessons } from '../services/api';

function LessonList() {
  const [lessons, setLessons] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchLessons = async () => {
      try {
        const response = await getLessons();
        setLessons(response.data);
        setLoading(false);
      } catch (err) {
        setError('Erro ao carregar lições');
        setLoading(false);
      }
    };

    fetchLessons();
  }, []);

  if (loading) return <div>Carregando...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className="lesson-list">
      <h2>Lições Disponíveis</h2>
      {lessons.map(lesson => (
        <div key={lesson.id} className="lesson-item">
          <h3>{lesson.title}</h3>
          <p>{lesson.description}</p>
          <span>Linguagem: {lesson.language}</span>
          <span>Dificuldade: {lesson.difficulty}</span>
        </div>
      ))}
    </div>
  );
}

export default LessonList;
