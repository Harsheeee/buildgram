import { useState } from 'react';
import { Link } from 'react-router-dom';
import { Heart, MessageCircle, Bookmark, MoreHorizontal, Send } from 'lucide-react';
import { Avatar } from '@/components/Avatar';
import { interactionAPI } from '@/services/api';
import { cn } from '@/lib/utils';
import { useAuth } from '@/contexts/AuthContext';

interface PostUser {
  id: number;
  username: string;
  profile_picture_url: string;
}

interface Comment {
  id: number;
  content: string;
  created_at: string;
  user: PostUser;
}

interface PostCardProps {
  id: number;
  user: PostUser;
  image_url: string;
  caption: string;
  like_count: number;
  comment_count: number;
  is_liked: boolean;
  created_at: string;
  comments?: Comment[];
}

export function PostCard({
  id,
  user: postUser,
  image_url,
  caption,
  like_count: initialLikeCount,
  comment_count: initialCommentCount,
  is_liked: initialIsLiked,
  created_at,
  comments: initialComments = [],
}: PostCardProps) {
  const { user: currentUser } = useAuth();
  const [isLiked, setIsLiked] = useState(initialIsLiked);
  const [likeCount, setLikeCount] = useState(initialLikeCount);
  const [commentCount, setCommentCount] = useState(initialCommentCount);
  const [comments, setComments] = useState<Comment[]>(initialComments);
  const [commentText, setCommentText] = useState('');
  const [isLikeAnimating, setIsLikeAnimating] = useState(false);
  const [showDoubleHeartAnim, setShowDoubleHeartAnim] = useState(false);
  const [isSubmittingComment, setIsSubmittingComment] = useState(false);

  const timeAgo = (dateStr: string) => {
    const seconds = Math.floor((Date.now() - new Date(dateStr).getTime()) / 1000);
    if (seconds < 60) return `${seconds}s`;
    const minutes = Math.floor(seconds / 60);
    if (minutes < 60) return `${minutes}m`;
    const hours = Math.floor(minutes / 60);
    if (hours < 24) return `${hours}h`;
    const days = Math.floor(hours / 24);
    if (days < 7) return `${days}d`;
    const weeks = Math.floor(days / 7);
    return `${weeks}w`;
  };

  const handleLike = async () => {
    setIsLikeAnimating(true);
    setTimeout(() => setIsLikeAnimating(false), 300);

    try {
      const res = await interactionAPI.toggleLike(id);
      setIsLiked(res.data.is_liked);
      setLikeCount(res.data.like_count);
    } catch (err) {
      console.error('Failed to toggle like:', err);
    }
  };

  const handleDoubleTap = async () => {
    if (!isLiked) {
      setShowDoubleHeartAnim(true);
      setTimeout(() => setShowDoubleHeartAnim(false), 1000);
      try {
        const res = await interactionAPI.toggleLike(id);
        setIsLiked(res.data.is_liked);
        setLikeCount(res.data.like_count);
      } catch (err) {
        console.error('Failed to like:', err);
      }
    } else {
      setShowDoubleHeartAnim(true);
      setTimeout(() => setShowDoubleHeartAnim(false), 1000);
    }
  };

  const handleComment = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!commentText.trim() || isSubmittingComment) return;

    setIsSubmittingComment(true);
    try {
      const res = await interactionAPI.addComment(id, commentText.trim());
      setComments([res.data, ...comments]);
      setCommentCount((prev) => prev + 1);
      setCommentText('');
    } catch (err) {
      console.error('Failed to add comment:', err);
    } finally {
      setIsSubmittingComment(false);
    }
  };

  return (
    <article className="bg-zinc-900/50 border border-zinc-800 rounded-2xl overflow-hidden backdrop-blur-sm transition-all duration-300 hover:border-zinc-700">
      {/* Header */}
      <div className="flex items-center justify-between px-4 py-3">
        <Link to={`/profile/${postUser.id}`} className="flex items-center gap-3 group">
          <Avatar src={postUser.profile_picture_url} alt={postUser.username} size="sm" />
          <div>
            <p className="text-sm font-semibold text-white group-hover:text-purple-400 transition-colors">
              {postUser.username}
            </p>
          </div>
        </Link>
        <button className="text-zinc-500 hover:text-white transition-colors p-1">
          <MoreHorizontal className="w-5 h-5" />
        </button>
      </div>

      {/* Image */}
      <div
        className="relative aspect-square bg-zinc-800 cursor-pointer"
        onDoubleClick={handleDoubleTap}
      >
        <img
          src={image_url}
          alt={caption || 'Post'}
          className="w-full h-full object-cover"
          loading="lazy"
        />
        {/* Double-tap heart animation */}
        {showDoubleHeartAnim && (
          <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
            <Heart
              className="w-24 h-24 text-white fill-white animate-[heartPop_0.8s_ease-out_forwards] drop-shadow-2xl"
            />
          </div>
        )}
      </div>

      {/* Actions */}
      <div className="px-4 pt-3">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <button
              onClick={handleLike}
              className={cn(
                'transition-all duration-200 hover:scale-110',
                isLikeAnimating && 'animate-[heartBounce_0.3s_ease-out]'
              )}
            >
              <Heart
                className={cn(
                  'w-6 h-6 transition-colors duration-200',
                  isLiked ? 'text-red-500 fill-red-500' : 'text-zinc-300 hover:text-zinc-100'
                )}
              />
            </button>
            <button className="text-zinc-300 hover:text-zinc-100 transition-colors hover:scale-110 duration-200">
              <MessageCircle className="w-6 h-6" />
            </button>
            <button className="text-zinc-300 hover:text-zinc-100 transition-colors hover:scale-110 duration-200">
              <Send className="w-6 h-6" />
            </button>
          </div>
          <button className="text-zinc-300 hover:text-zinc-100 transition-colors hover:scale-110 duration-200">
            <Bookmark className="w-6 h-6" />
          </button>
        </div>

        {/* Likes */}
        <p className="text-sm font-semibold text-white mt-3">
          {likeCount.toLocaleString()} {likeCount === 1 ? 'like' : 'likes'}
        </p>

        {/* Caption */}
        {caption && (
          <p className="text-sm text-zinc-300 mt-1">
            <Link to={`/profile/${postUser.id}`} className="font-semibold text-white hover:text-purple-400 transition-colors mr-1.5">
              {postUser.username}
            </Link>
            {caption}
          </p>
        )}

        {/* Comment count */}
        {commentCount > 0 && (
          <Link to={`/post/${id}`} className="text-sm text-zinc-500 mt-1 block hover:text-zinc-400 transition-colors">
            View all {commentCount} comments
          </Link>
        )}

        {/* Preview comments */}
        {comments.slice(0, 2).map((comment) => (
          <p key={comment.id} className="text-sm text-zinc-300 mt-1">
            <Link to={`/profile/${comment.user.id}`} className="font-semibold text-white mr-1.5 hover:text-purple-400 transition-colors">
              {comment.user.username}
            </Link>
            {comment.content}
          </p>
        ))}

        {/* Timestamp */}
        <p className="text-xs text-zinc-600 mt-2 uppercase tracking-wider">
          {timeAgo(created_at)}
        </p>
      </div>

      {/* Comment Input */}
      <form onSubmit={handleComment} className="flex items-center gap-3 px-4 py-3 mt-1 border-t border-zinc-800/50">
        <Avatar src={currentUser?.profile_picture_url} alt={currentUser?.username || 'U'} size="sm" />
        <input
          type="text"
          placeholder="Add a comment..."
          value={commentText}
          onChange={(e) => setCommentText(e.target.value)}
          className="flex-1 bg-transparent text-sm text-zinc-300 placeholder:text-zinc-600 focus:outline-none"
        />
        {commentText.trim() && (
          <button
            type="submit"
            disabled={isSubmittingComment}
            className="text-sm font-semibold text-purple-400 hover:text-purple-300 transition-colors disabled:opacity-50"
          >
            Post
          </button>
        )}
      </form>
    </article>
  );
}
