//ReactDOM.render( <h1>Hello, world!</h1>, document.getElementById('root') )
const LikeList = props => {
    return <ul className="item-list">
    {props.likes.map(item => <LikeDetail item={item} />)}
    </ul>;
};

const LikeDetail = props => {
    const item = props.item;
    return <li>
        <div className="item-time">{item.Timestamp}</div>
    </li>;
};

const CommentList = props => {
    return <ul className="item-list">
    {props.comments.map(item => <CommentDetail item={item} />)}
    </ul>;
};

const CommentDetail = props => {
    const item = props.item;
    return <li>
        <div className="item-time">{item.Timestamp}</div>
        <div className="item-time">{item.Content}</div>
    </li>;
};



fetch('/api/likes').then(res => res.json()).then(data => {
    ReactDOM.render(
        <LikeList likes={data.Likes} />, // これを
        document.getElementById('likes-board') // ここにレンダリングしろ
    );
});

fetch('api/comments').then(res => res.json()).then(data => {
    ReactDOM.render(
        <CommentList comments={data.Comments} />,
        document.getElementById('comments-board')
    );
});

document.getElementById('like').onclick = function() {
    fetch('api/likes', {
        method : 'POST',
        body : '',
    }).then(res => res.text())
    .then(text => console.log(text));
};

document.getElementById('comment').onclick = function() {
    const content = document.getElementById("content").value;
    if (content.length > 0) {
        fetch('api/comments', {
            method : 'POST',
            body : 'content=' + content,
            headers : new Headers({'Content-type' : 'application/x-www-form-urlencoded' })
        }).then(res => res.text())
        .then(text => console.log(text));
    }
};