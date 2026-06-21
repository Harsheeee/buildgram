import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { postAPI, userAPI } from '@/services/api';
import { Input } from '@/components/Input';
import { Avatar } from '@/components/Avatar';
import { Search, Loader2 } from 'lucide-react';

export default function ExplorePage() {
  const [posts, setPosts] = useState<any[]>([]);
  const [searchQuery, setSearchQuery] = useState('');
  const [searchResults, setSearchResults] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isSearching, setIsSearching] = useState(false);

  useEffect(() => { loadExplorePosts(); }, []);

  const loadExplorePosts = async () => {
    try {
      const res = await postAPI.getExplorePosts(1, 30);
      setPosts(res.data.posts || []);
    } catch (err) { console.error(err); }
    finally { setIsLoading(false); }
  };

  useEffect(() => {
    if (!searchQuery.trim()) { setSearchResults([]); return; }
    const timer = setTimeout(async () => {
      setIsSearching(true);
      try {
        const res = await userAPI.searchUsers(searchQuery);
        setSearchResults(res.data || []);
      } catch (err) { console.error(err); }
      finally { setIsSearching(false); }
    }, 300);
    return () => clearTimeout(timer);
  }, [searchQuery]);

  if (isLoading) {
    return (<div className="flex items-center justify-center min-h-[60vh]"><Loader2 className="w-8 h-8 text-purple-500 animate-spin" /></div>);
  }

  return (
    <div className="max-w-4xl mx-auto px-4 py-6">
      <div className="mb-6">
        <Input id="explore-search" placeholder="Search users..." value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)} icon={<Search className="w-4 h-4" />} />
      </div>
      {searchQuery.trim() ? (
        <div className="space-y-2">
          {isSearching ? (<div className="flex justify-center py-8"><Loader2 className="w-6 h-6 text-purple-500 animate-spin" /></div>)
            : searchResults.length === 0 ? (<p className="text-center text-zinc-500 py-8">No users found</p>)
            : searchResults.map((user: any) => (
              <Link key={user.id} to={`/profile/${user.id}`}
                className="flex items-center gap-3 p-3 rounded-xl hover:bg-zinc-800/50 transition-all">
                <Avatar src={user.profile_picture_url} alt={user.username} size="md" />
                <div><p className="text-sm font-semibold text-white">{user.username}</p>
                  <p className="text-xs text-zinc-500">{user.full_name}</p></div>
              </Link>))}
        </div>
      ) : (
        <div className="grid grid-cols-3 gap-1">
          {posts.map((post: any) => (
            <Link key={post.id} to={`/post/${post.id}`}
              className="relative aspect-square group overflow-hidden bg-zinc-800">
              <img src={post.image_url} alt="" className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-105" loading="lazy" />
              <div className="absolute inset-0 bg-black/0 group-hover:bg-black/40 transition-all duration-300 flex items-center justify-center opacity-0 group-hover:opacity-100">
                <div className="flex items-center gap-6 text-white font-semibold">
                  <span>❤️ {post.like_count}</span><span>💬 {post.comment_count}</span>
                </div>
              </div>
            </Link>
          ))}
        </div>
      )}
    </div>
  );
}
