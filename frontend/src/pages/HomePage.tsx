import { useState, useEffect } from 'react';
import { PostCard } from '@/components/PostCard';
import { postAPI } from '@/services/api';
import { Loader2 } from 'lucide-react';

export default function HomePage() {
  const [posts, setPosts] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);

  const loadPosts = async (pageNum: number) => {
    try {
      const res = await postAPI.getFeed(pageNum);
      const newPosts = res.data.posts || [];
      if (pageNum === 1) {
        setPosts(newPosts);
      } else {
        setPosts((prev) => [...prev, ...newPosts]);
      }
      setHasMore(newPosts.length >= 10);
    } catch (err) {
      console.error('Failed to load feed:', err);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    loadPosts(1);
  }, []);

  const loadMore = () => {
    const nextPage = page + 1;
    setPage(nextPage);
    loadPosts(nextPage);
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-[60vh]">
        <Loader2 className="w-8 h-8 text-purple-500 animate-spin" />
      </div>
    );
  }

  return (
    <div className="max-w-lg mx-auto px-4 py-6 space-y-6">
      {posts.length === 0 ? (
        <div className="text-center py-20">
          <div className="w-20 h-20 mx-auto mb-6 rounded-full bg-zinc-800/50 flex items-center justify-center">
            <span className="text-4xl">📷</span>
          </div>
          <h2 className="text-xl font-semibold text-white mb-2">Your feed is empty</h2>
          <p className="text-zinc-500 text-sm max-w-xs mx-auto">
            Follow other users or create a post to see content here.
          </p>
        </div>
      ) : (
        <>
          {posts.map((post) => (
            <PostCard key={post.id} {...post} />
          ))}
          {hasMore && (
            <button
              onClick={loadMore}
              className="w-full py-3 text-sm text-purple-400 hover:text-purple-300 transition-colors"
            >
              Load more posts
            </button>
          )}
        </>
      )}
    </div>
  );
}
