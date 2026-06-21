import { useState, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import { postAPI } from '@/services/api';
import { Button } from '@/components/Button';
import { ImagePlus, X } from 'lucide-react';

export default function CreatePostPage() {
  const [file, setFile] = useState<File | null>(null);
  const [preview, setPreview] = useState<string | null>(null);
  const [caption, setCaption] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState('');
  const fileInputRef = useRef<HTMLInputElement>(null);
  const navigate = useNavigate();

  const handleFileSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0];
    if (!selectedFile) return;
    const validTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp'];
    if (!validTypes.includes(selectedFile.type)) {
      setError('Please select a valid image file (JPG, PNG, GIF, or WebP)');
      return;
    }
    if (selectedFile.size > 10 * 1024 * 1024) {
      setError('Image must be less than 10MB');
      return;
    }
    setFile(selectedFile);
    setError('');
    const reader = new FileReader();
    reader.onload = (ev) => setPreview(ev.target?.result as string);
    reader.readAsDataURL(selectedFile);
  };

  const clearFile = () => {
    setFile(null);
    setPreview(null);
    if (fileInputRef.current) fileInputRef.current.value = '';
  };

  const handleSubmit = async () => {
    if (!file) { setError('Please select an image'); return; }
    setIsLoading(true);
    setError('');
    try {
      await postAPI.createPost(file, caption);
      navigate('/');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to create post');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="max-w-lg mx-auto px-4 py-6">
      <h1 className="text-xl font-bold text-white mb-6">Create new post</h1>
      {error && (
        <div className="bg-red-500/10 border border-red-500/20 text-red-400 text-sm rounded-lg px-4 py-3 mb-4">{error}</div>
      )}
      {!preview ? (
        <div onClick={() => fileInputRef.current?.click()}
          className="aspect-square rounded-2xl border-2 border-dashed border-zinc-700 bg-zinc-900/50 flex flex-col items-center justify-center cursor-pointer hover:border-purple-500/50 hover:bg-zinc-800/50 transition-all duration-300 group">
          <div className="w-16 h-16 rounded-full bg-zinc-800 group-hover:bg-purple-500/10 flex items-center justify-center mb-4 transition-colors">
            <ImagePlus className="w-8 h-8 text-zinc-500 group-hover:text-purple-400 transition-colors" />
          </div>
          <p className="text-zinc-400 text-sm font-medium">Click to upload a photo</p>
          <p className="text-zinc-600 text-xs mt-1">JPG, PNG, GIF or WebP (max 10MB)</p>
        </div>
      ) : (
        <div className="relative aspect-square rounded-2xl overflow-hidden bg-zinc-800">
          <img src={preview} alt="Preview" className="w-full h-full object-cover" />
          <button onClick={clearFile} className="absolute top-3 right-3 w-8 h-8 rounded-full bg-black/70 backdrop-blur-sm flex items-center justify-center text-white hover:bg-black/90 transition-colors">
            <X className="w-4 h-4" />
          </button>
        </div>
      )}
      <input ref={fileInputRef} type="file" accept="image/*" onChange={handleFileSelect} className="hidden" />
      <div className="mt-4">
        <textarea placeholder="Write a caption..." value={caption} onChange={(e) => setCaption(e.target.value)}
          maxLength={2200} rows={4}
          className="w-full bg-zinc-900/50 border border-zinc-700 rounded-xl px-4 py-3 text-sm text-white placeholder:text-zinc-500 focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-500 resize-none transition-all" />
        <p className="text-right text-xs text-zinc-600 mt-1">{caption.length}/2200</p>
      </div>
      <Button onClick={handleSubmit} className="w-full mt-4" isLoading={isLoading} disabled={!file}>Share Post</Button>
    </div>
  );
}
