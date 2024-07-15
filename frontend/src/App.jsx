import { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [books, setBooks] = useState([]);
  const [title, setTitle] = useState('');
  const [author, setAuthor] = useState('');
  const [editing, setEditing] = useState(false);
  const [currentBook, setCurrentBook] = useState({});

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        console.log('API URL:', import.meta.env.VITE_API_URL); 
        const response = await axios.get(`${import.meta.env.VITE_API_URL}/books`);
        setBooks(response.data);
      } catch (error) {
        console.error('Error fetching books:', error);
      }
    };
    fetchBooks();
  }, []);

  const addBook = async () => {
    try {
      const response = await axios.post(`${import.meta.env.VITE_API_URL}/books`, {
        title,
        author
      });
      setBooks([...books, response.data]);
    } catch (error) {
      console.error('Error adding book', error);
    }
  };

  const updateBook = async () => {
    try {
      await axios.put(`${import.meta.env.VITE_API_URL}/books/${currentBook.id}`, {
        title,
        author 
      });
      setBooks(books.map(book => (book.id === currentBook.id ? { ...book, title, author } : book)));
      setEditing(false);
      setTitle('');
      setAuthor('');
    } catch (error) {
      console.error('Error updating book:', error);
    }
  };

  const deleteBook = async (id) => {
    try {
      await axios.delete(`${import.meta.env.VITE_API_URL}/books/${id}`);
      setBooks(books.filter(book => book.id !== id));
    } catch (error) {
      console.error('Error deleting book:', error);
    }
  };

  const editBook = (book) => {
    setEditing(true);
    setCurrentBook(book);
    setTitle(book.title);
    setAuthor(book.author);
  };

  return (
    <div className='App'>
      <h1>Books</h1>
      <ul>
        {books.map(book => (
          <li key={book.id}>
            {book.title} by {book.author}
            <button onClick={() => editBook(book)}>Edit</button>
            <button onClick={() => deleteBook(book.id)}>Delete</button>
            </li>
        ))}
      </ul>
      <h2>{editing ? 'Edit Book' : 'Add Book'}</h2>
      <input
        type='text'
        placeholder='Title'
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />
      <input
        type='text'
        placeholder='Author'
        value={author}
        onChange={(e) => setAuthor(e.target.value)}
      />
      <button onClick={editing ? updateBook : addBook}>
        {editing ? 'Update' : 'Add'}
      </button>
    </div>
  )
}

export default App
