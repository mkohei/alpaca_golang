//ReactDOM.render( <h1>Hello, world!</h1>, document.getElementById('root') )
const ItemList = props => {
    return <ul className="item-list">
    {props.likes.map(item => <ItemDetail item={item} />)}
    </ul>;
};

const ItemDetail = props => {
    const item = props.item;
    return <li>
        <div className="item-time">{item.Timestamp}</div>
    </li>;
};


fetch('/api/likes').then(res => res.json()).then(data => {
    ReactDOM.render(
        <ItemList likes={data.Likes} />, // これを
        document.getElementById('root') // ここにレンダリングしろ
    );
});

document.getElementById('like').onclick = function() {
    //
    fetch('api/likes', {
        method : 'POST',
        body : '',
    }).then(res => res.text())
    .then(text => console.log(text));
};