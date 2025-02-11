import NavBar from './NavBar';
import { useState, useEffect } from 'react';
import { SelectedPage } from '@/Components/Shared/Types';

const Header = () => {
  const [selectedPage, setSelectedPage] = useState<SelectedPage>(
    SelectedPage.Home
  );
  // const [isTopOfPage setIsTopOfPage] = useState<boolean>(true);
  const flexBetween = 'flex items-center justify-between';

  useEffect(() => {
    const handleScroll = () => {
      
      if (window.scrollY === 0) {
        setSelectedPage(SelectedPage.Home);
      }
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  return (
    <div className={`${flexBetween} bg-white transition fixed top-0 z-30 w-full p-5 md:px-16`}>
      <h1 className="text-xl sm:text-2xl font-bold text-green-500">Health Mentá</h1>
      <NavBar
        flexBetween={flexBetween}
        selectedPage={selectedPage}
        setSelectedPage={setSelectedPage}
      />
    </div>
  );
};

export default Header;
