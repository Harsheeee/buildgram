import { useState } from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { useAuth } from '@/contexts/AuthContext';
import { Avatar } from '@/components/Avatar';
import { Home, Search, PlusSquare, Heart, User, LogOut, Menu, X } from 'lucide-react';
import { cn } from '@/lib/utils';

export function Navbar() {
  const { user, logout } = useAuth();
  const location = useLocation();
  const navigate = useNavigate();
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  const navItems = [
    { path: '/', icon: Home, label: 'Home' },
    { path: '/explore', icon: Search, label: 'Explore' },
    { path: '/create', icon: PlusSquare, label: 'Create' },
    { path: '/activity', icon: Heart, label: 'Activity' },
    { path: `/profile/${user?.id}`, icon: User, label: 'Profile' },
  ];

  return (
    <>
      {/* Desktop Sidebar */}
      <nav className="hidden md:flex fixed left-0 top-0 h-full w-[72px] lg:w-[240px] flex-col border-r border-zinc-800 bg-black/95 backdrop-blur-xl z-50 transition-all duration-300">
        {/* Logo */}
        <Link to="/" className="flex items-center gap-3 px-4 py-6 lg:px-6">
          <div className="w-8 h-8 rounded-lg bg-gradient-to-br from-pink-500 via-purple-500 to-indigo-500 flex items-center justify-center flex-shrink-0">
            <span className="text-white font-bold text-sm">B</span>
          </div>
          <span className="hidden lg:block text-xl font-bold bg-gradient-to-r from-pink-400 via-purple-400 to-indigo-400 bg-clip-text text-transparent">
            BuildGram
          </span>
        </Link>

        {/* Nav Items */}
        <div className="flex-1 flex flex-col gap-1 px-3 mt-4">
          {navItems.map((item) => {
            const isActive = location.pathname === item.path;
            return (
              <Link
                key={item.path}
                to={item.path}
                className={cn(
                  'flex items-center gap-4 px-3 py-3 rounded-xl transition-all duration-200 group',
                  isActive
                    ? 'bg-zinc-800/80 text-white'
                    : 'text-zinc-400 hover:text-white hover:bg-zinc-800/50'
                )}
              >
                <item.icon
                  className={cn(
                    'w-6 h-6 flex-shrink-0 transition-transform duration-200',
                    isActive ? 'scale-105' : 'group-hover:scale-105'
                  )}
                  strokeWidth={isActive ? 2.5 : 1.5}
                />
                <span className={cn('hidden lg:block text-sm', isActive && 'font-semibold')}>
                  {item.label}
                </span>
              </Link>
            );
          })}
        </div>

        {/* User section */}
        <div className="px-3 pb-6">
          <div className="flex items-center gap-3 px-3 py-3 rounded-xl hover:bg-zinc-800/50 transition-all duration-200 cursor-pointer"
            onClick={() => navigate(`/profile/${user?.id}`)}
          >
            <Avatar src={user?.profile_picture_url} alt={user?.username || 'U'} size="sm" />
            <div className="hidden lg:block flex-1 min-w-0">
              <p className="text-sm font-semibold text-white truncate">{user?.username}</p>
              <p className="text-xs text-zinc-500 truncate">{user?.full_name}</p>
            </div>
          </div>
          <button
            onClick={logout}
            className="flex items-center gap-4 px-3 py-3 rounded-xl text-zinc-400 hover:text-red-400 hover:bg-zinc-800/50 transition-all duration-200 w-full mt-1"
          >
            <LogOut className="w-6 h-6 flex-shrink-0" strokeWidth={1.5} />
            <span className="hidden lg:block text-sm">Log out</span>
          </button>
        </div>
      </nav>

      {/* Mobile Bottom Nav */}
      <nav className="md:hidden fixed bottom-0 left-0 right-0 h-16 bg-black/95 backdrop-blur-xl border-t border-zinc-800 z-50 flex items-center justify-around px-2">
        {navItems.map((item) => {
          const isActive = location.pathname === item.path;
          return (
            <Link
              key={item.path}
              to={item.path}
              className={cn(
                'flex flex-col items-center gap-0.5 p-2 rounded-xl transition-all duration-200',
                isActive ? 'text-white' : 'text-zinc-500'
              )}
            >
              <item.icon className="w-6 h-6" strokeWidth={isActive ? 2.5 : 1.5} />
            </Link>
          );
        })}
      </nav>

      {/* Mobile Top Bar */}
      <div className="md:hidden fixed top-0 left-0 right-0 h-14 bg-black/95 backdrop-blur-xl border-b border-zinc-800 z-50 flex items-center justify-between px-4">
        <span className="text-lg font-bold bg-gradient-to-r from-pink-400 via-purple-400 to-indigo-400 bg-clip-text text-transparent">
          BuildGram
        </span>
        <button
          onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
          className="text-zinc-400 hover:text-white"
        >
          {mobileMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
        </button>
      </div>

      {/* Mobile Menu Overlay */}
      {mobileMenuOpen && (
        <div className="md:hidden fixed inset-0 bg-black/90 backdrop-blur-sm z-40 pt-14 pb-16">
          <div className="flex flex-col items-center justify-center h-full gap-2">
            <div className="flex items-center gap-3 mb-6">
              <Avatar src={user?.profile_picture_url} alt={user?.username || 'U'} size="lg" />
              <div>
                <p className="text-lg font-bold text-white">{user?.username}</p>
                <p className="text-sm text-zinc-500">{user?.full_name}</p>
              </div>
            </div>
            <button
              onClick={() => { logout(); setMobileMenuOpen(false); }}
              className="flex items-center gap-3 px-6 py-3 rounded-xl text-red-400 hover:bg-zinc-800/50 transition-all"
            >
              <LogOut className="w-5 h-5" />
              <span>Log out</span>
            </button>
          </div>
        </div>
      )}
    </>
  );
}
