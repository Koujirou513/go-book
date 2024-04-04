// APIのエンドポイントを定義
const apiUrl = 'http://localhost:8080/books';

// 本のデータをAPIから取得して表示する関数
async function fetchBooks() {
    const response = await fetch(apiUrl); // APIからデータを取得するためにGETリクエストを送信する
    const books = await response.json(); // レスポンスのボディをJSON形式で解析
    const booksContainer = document.getElementById('books'); // 本を表示するコンテナを取得
    booksContainer.innerHTML = ''; // コンテナの内容を初期化
    books.forEach(book => {  // 取得した本のデータを1つずつ処理する
        const bookDiv = document.createElement('div'); // 本の情報を表示するための新しいdiv要素を作成
        bookDiv.className = 'book'; // 作成したdivにクラス名を設定
        // 本のタイトル、著者、更新ボタン、削除ボタンをdivに追加する
        bookDiv.innerHTML = `   
            <div>Title: ${book.title}</div>
            <div>Author: ${book.author}</div>
            <div>
                <button onclick="updateBook(${book.id})">Update</button>
                <button onclick="deleteBook(${book.id})">Delete</button>
            </div>
        `;
        booksContainer.appendChild(bookDiv) // 完成したdivを本を表示するコンテナに追加する
    })
}

// 新しい本を追加する関数を定義
async function createBook() {
    const title = document.getElementById('newTitle').value; // タイトルを入力フィールドから取得
    const author = document.getElementById('newAuthor').value; // 著者名を入力フィールドから取得
    // 入力されたタイトルと著者名で新しい本をAPIにPOSTリクエストを送信して追加する
    await fetch(apiUrl, {
        method: 'POST', 
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, author }) // 本のデータをJSON形式でリクエストボディに設定
    });
    fetchBooks(); // 本のリストを更新
    // インプットフィールドをクリア
    document.getElementById('newTitle').value = '';
    document.getElementById('newAuthor').value = '';
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