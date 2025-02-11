import { Bars3Icon, XMarkIcon } from '@heroicons/react/20/solid';
import { SelectedPage } from '@/Components/Shared/Types';
import { useState } from 'react';
import { useNavigate,Link } from 'react-router-dom';
import useMediaQuery from '@/Hooks/useMediaQuery';

// import Links from './Links';
import Button from '../UI/Button';

type Props = {
  flexBetween: string;
  selectedPage: SelectedPage;
  setSelectedPage: (value: SelectedPage) => void;
};

const NavBar = ({ flexBetween }: Props) => {
  const [isMenuToggled, setIsMenuToggled] = useState<boolean>(false);
  const isAboveMediumScreens = useMediaQuery('(min-width: 900px)');
  const navigate = useNavigate();

  const handleHomeClick = () => {
    navigate('/Hospital-Website'); // Navigate to Services page
  };
  
  const handleDoctorsClick = () => {
    navigate('/Hospital-Website/doctors'); // Navigate to Doctors page
  };

  const handleAboutClick = () => {
    navigate('/Hospital-Website/about'); // Navigate to About page
  };

  const handleServicesClick = () => {
    navigate('/Hospital-Website/services'); // Navigate to Services page
  };

  return (
    <nav>
      {/* Desktop Menu */}
      {isAboveMediumScreens && (
        <div className={`${flexBetween} lg:gap-28 gap-20`}>
          <div className={`${flexBetween} gap-16`}>
            {/* <Links selectedPage={selectedPage} setSelectedPage={setSelectedPage} /> */}

            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleHomeClick}
            >
              Home
            </button>
            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleDoctorsClick}
            >
              Doctors
            </button>
            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleAboutClick}
            >
              About
            </button>
            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleServicesClick}
            >
              Services
            </button>
          </div>
          <Button >
            <Link to="/Hospital-Website/login">Login</Link>
          </Button>
        </div>
      )}

      {/* Mobile Menu Toggle Button */}
      {!isAboveMediumScreens && (
        <button onClick={() => setIsMenuToggled((prev) => !prev)}>
          <Bars3Icon className="h-8 w-8" />
        </button>
      )}

      {/* Mobile Menu Modal */}
      {!isAboveMediumScreens && isMenuToggled && (
        <div className="fixed right-0 top-0 z-40 h-80 rounded-es-3xl w-[175px] md:w-[300px] bg-secondary drop-shadow-2xl">
          {/* Close Button */}
          <div className="flex justify-end p-5 md:pr-16 sm:pt-10">
            <button onClick={() => setIsMenuToggled((prev) => !prev)}>
              <XMarkIcon className="h-10 w-10" />
            </button>
          </div>

          {/* Mobile Menu Items */}
          <div className="ml-[20%] flex flex-col items-start gap-5 text-2xl">
            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleHomeClick}
            >
              Home
            </button>
            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleDoctorsClick}
            >
              Doctors
            </button>
            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleAboutClick}
            >
              About
            </button>
            <button
              className="text-lg font-bold text-primary hover:text-[#2b7dad] transition duration-500"
              onClick={handleServicesClick}
            >
              Services
            </button>
          </div>
        </div>
      )}
    </nav>
  );
};

export default NavBar;
