import { useState, useRef } from 'react';
import { useAuth } from '@/contexts/AuthContext';
import { userAPI } from '@/services/api';
import { Button } from '@/components/Button';
import { Input } from '@/components/Input';
import { Avatar } from '@/components/Avatar';
import { Camera } from 'lucide-react';

export default function SettingsPage() {
  const { user, updateUser } = useAuth();
  const [fullName, setFullName] = useState(user?.full_name || '');
  const [username, setUsername] = useState(user?.username || '');
  const [bio, setBio] = useState(user?.bio || '');
  const [isLoading, setIsLoading] = useState(false);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleSave = async () => {
    setIsLoading(true); setError(''); setMessage('');
    try {
      const res = await userAPI.updateProfile({ full_name: fullName, username, bio });
      updateUser({ ...user!, ...res.data });
      setMessage('Profile updated successfully!');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to update profile');
    } finally { setIsLoading(false); }
  };

  const handlePictureUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    try {
      const res = await userAPI.uploadProfilePicture(file);
      updateUser({ ...user!, ...res.data });
      setMessage('Profile picture updated!');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Failed to upload picture');
    }
  };

  return (
    <div className="max-w-lg mx-auto px-4 py-6">
      <h1 className="text-xl font-bold text-white mb-6">Edit Profile</h1>
      {message && <div className="bg-green-500/10 border border-green-500/20 text-green-400 text-sm rounded-lg px-4 py-3 mb-4">{message}</div>}
      {error && <div className="bg-red-500/10 border border-red-500/20 text-red-400 text-sm rounded-lg px-4 py-3 mb-4">{error}</div>}

      <div className="bg-zinc-900/50 border border-zinc-800 rounded-2xl p-6 mb-6">
        <div className="flex items-center gap-4">
          <div className="relative">
            <Avatar src={user?.profile_picture_url} alt={user?.username || 'U'} size="lg" />
            <button onClick={() => fileInputRef.current?.click()}
              className="absolute -bottom-1 -right-1 w-7 h-7 rounded-full bg-purple-500 flex items-center justify-center text-white hover:bg-purple-400 transition-colors">
              <Camera className="w-3.5 h-3.5" />
            </button>
          </div>
          <div>
            <p className="font-semibold text-white">{user?.username}</p>
            <button onClick={() => fileInputRef.current?.click()} className="text-sm text-purple-400 hover:text-purple-300">
              Change profile photo
            </button>
          </div>
        </div>
        <input ref={fileInputRef} type="file" accept="image/*" onChange={handlePictureUpload} className="hidden" />
      </div>

      <div className="space-y-4">
        <Input id="settings-fullname" label="Full Name" value={fullName} onChange={(e) => setFullName(e.target.value)} />
        <Input id="settings-username" label="Username" value={username} onChange={(e) => setUsername(e.target.value)} />
        <div className="space-y-1.5">
          <label className="block text-sm font-medium text-zinc-300">Bio</label>
          <textarea value={bio} onChange={(e) => setBio(e.target.value)} maxLength={500} rows={3}
            className="w-full bg-zinc-900/50 border border-zinc-700 rounded-lg px-4 py-2 text-sm text-white placeholder:text-zinc-500 focus:outline-none focus:ring-2 focus:ring-purple-500/50 focus:border-purple-500 resize-none transition-all" />
        </div>
        <Button onClick={handleSave} className="w-full" isLoading={isLoading}>Save Changes</Button>
      </div>
    </div>
  );
}
