import { useParams } from "react-router-dom"
export default function SinglePost() {
    const { id } = useParams();
    return (
        <div className="postDetails">
            <h2>{id}</h2>
        </div>
    )
}
