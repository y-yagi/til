use notify::{Config, RecommendedWatcher, RecursiveMode, Watcher};
use rutie::{class, methods, Boolean, Class, NilClass, Object, RString, Thread, VM};
use std::path::Path;
use std::sync::mpsc;

class!(RutieExample);
class!(Executor);

methods! {
    RutieExample,
    _rtself,

    fn watch(input: RString) -> NilClass {
        let ruby_string = input.map_err(|e| VM::raise_ex(e) ).unwrap();
        if !VM::is_block_given() {
            VM::raise(Class::from_existing("LocalJumpError"), "no block given (yield)");
            unreachable!();
        }

        Thread::new(move || {
            let path = Path::new(ruby_string.to_str());
            let (tx, rx) = mpsc::channel();
            let mut watcher = RecommendedWatcher::new(tx, Config::default()).map_err(|_e| VM::raise(Class::from_existing("RuntimeError"), "watch starting failed")).unwrap();

            watcher.watch(path, RecursiveMode::Recursive).map_err(|_e| VM::raise(Class::from_existing("RuntimeError"), "watch starting failed")).unwrap();

            let handler = move || {
                loop {
                    let res = rx.recv();
                    match res {
                        Ok(event) => println!("changed: {:?}", event),
                        Err(e) => println!("watch error: {:?}", e),
                    }
                }
            };

            Thread::call_without_gvl(handler, Some(|| {}));
            NilClass::new()
        });

        return NilClass::new()
    }
}

impl Executor {}

methods! {
    Executor,
    _rtself,

    fn stop() -> Boolean {
        return Boolean::new(true)
    }
}

#[allow(non_snake_case)]
#[no_mangle]
pub extern "C" fn Init_rutie_gvl_example() {
    Class::new("RutieExample", None).define(|klass| {
        klass.def_self("raw_watch", watch);
    });

    Class::new("Executor", None).define(|klass| {
        klass.def("stop", stop);
    });
}
