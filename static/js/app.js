// APIのエンドポイント
const apiUrl = 'http://localhost:8080/books';

// 本のデータをフェッチして表示する関数
async function fetchBooks() {
    const response = await fetch(apiUrl); // APIからデータをフェッチ
    const books = await response.json(); // レスポンスをJSON形式で解析
    const booksContainer = document.getElementById('books'); // 本を表示するコンテナを取得
    booksContainer.innerHTML = ''; // コンテナを初期化
    books.forEach(book => {
        const bookDiv = document.createElement('div'); // 各本のためのdivを作成
        bookDiv.className = 'book';
        bookDiv.innerHTML = `
            <div>Title: ${book.title}</div>
            <div>Author: ${book.author}</div>
            <div>
                <button onclick="updateBook(${book.id})">Update</button>
                <button onclick="deleteBook(${book.id})">Delete</button>
            </div>
        `;
        booksContainer.appendChild(bookDiv) // コンテナに本のdivを作成
    })
}

// 新しい本を追加する関数
async function createBook() {
    const title = document.getElementById('newTitle').value; // タイトルを入力フィールドから取得
    const author = document.getElementById('newAuthor').value; // 著者名を入力フィールドから取得
    await fetch(apiUrl, {
        method: 'POST', 
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, author }) // 本のデータをJSON形式でリクエストボディに設定
    });
    fetchBooks(); // 本のリストを更新
}

// 本のデータを更新する関数
async function updateBook(id) {
    const title = prompt('Enter new title:');
    const author = prompt('Enter new author');
    await fetch(`${apiUrl}/${id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, author })
    });
    fetchBooks();
}

// 本を削除する関数
async function deleteBook(id) {
    if (confirm('Are you sure you want to delete this book?')) { // 削除の確認
        await fetch(`${apiUrl}/${id}`, { method: 'DELETE' });
        fetchBooks();
    }
}

// 最初の本のデータ取得
fetchBooks();