import { cn } from '@/lib/utils';

interface AvatarProps {
  src?: string;
  alt?: string;
  size?: 'sm' | 'md' | 'lg' | 'xl';
  className?: string;
  onClick?: () => void;
}

const sizeMap = {
  sm: 'w-8 h-8',
  md: 'w-10 h-10',
  lg: 'w-16 h-16',
  xl: 'w-24 h-24',
};

export function Avatar({ src, alt = 'User', size = 'md', className, onClick }: AvatarProps) {
  return (
    <div
      className={cn(
        'relative rounded-full overflow-hidden bg-gradient-to-br from-pink-500 via-purple-500 to-indigo-500 p-[2px] flex-shrink-0',
        sizeMap[size],
        onClick && 'cursor-pointer hover:scale-105 transition-transform duration-200',
        className
      )}
      onClick={onClick}
    >
      <div className="w-full h-full rounded-full overflow-hidden bg-zinc-900">
        {src ? (
          <img
            src={src}
            alt={alt}
            className="w-full h-full object-cover"
          />
        ) : (
          <div className="w-full h-full flex items-center justify-center bg-zinc-800 text-zinc-400 font-semibold text-lg">
            {alt.charAt(0).toUpperCase()}
          </div>
        )}
      </div>
    </div>
  );
}
